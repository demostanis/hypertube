package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"maps"
	"os"
	"strconv"
	"sync"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/webtor-io/go-jackett"
)

const (
	jackettConfigPath = "./jackett/Jackett/ServerConfig.json"
	jackettURL        = "http://jackett:9117"
)

var defaultParams = map[string]string{
	"language": "fr-FR",
}

type TMDBClient struct {
	c    *tmdb.Client
	j    *jackett.Jackett
	log  *slog.Logger
	page int
}

func newTMDBClient(j *jackett.Jackett, logger *slog.Logger) (*TMDBClient, error) {
	client, err := tmdb.Init(os.Getenv("TMDB_API_KEY"))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize TMDB client: %w", err)
	}
	client.SetClientAutoRetry()
	return &TMDBClient{client, j, logger, 0}, nil
}

type Movie struct {
	Title string
}

func (t *TMDBClient) Jackettize(movie Movie) {
	t.log.Info("fetching movie", "movie", movie.Title)
	res, err := t.j.Fetch(context.TODO(), &jackett.FetchRequest{
		Query: movie.Title,
	})
	if err != nil {
		t.log.Error("failed to fetch movie",
			"movie", movie.Title,
			"error", err)
		return
	}

	if len(res.Results) == 0 {
		t.log.Warn("no torrents available for movie",
			"movie", movie.Title)
		return
	}

	t.log.Info("found torrents",
		"movie", movie.Title, "count", len(res.Results))
}

func (t *TMDBClient) Discover() error {
	t.page += 1
	params := maps.Clone(defaultParams)
	params["page"] = strconv.Itoa(t.page)

	movies, err := t.c.GetDiscoverMovie(params)
	if err != nil {
		return fmt.Errorf("failed to discover movies using TMDB: %w", err)
	}

	var wg sync.WaitGroup
	for _, movie := range movies.Results {
		wg.Add(1)
		go func() {
			t.Jackettize(Movie{movie.Title})
			wg.Done()
		}()
	}
	wg.Wait()

	return nil
}

func parseAPIKey() (string, error) {
	var apiKey struct {
		APIKey string
	}

	f, err := os.Open(jackettConfigPath)
	if err != nil {
		return "", fmt.Errorf("failed to open jackett config file at %s: %w",
			jackettConfigPath, err)
	}
	err = json.NewDecoder(f).Decode(&apiKey)
	if err != nil {
		return "", fmt.Errorf("failed to parse jackett config file: %w", err)
	}

	return apiKey.APIKey, nil
}

func run() error {
	f, _ := os.Create("/var/log/worker/worker.log")
	logger := slog.New(slog.NewJSONHandler(
		io.MultiWriter(os.Stdout, f),
		nil))

	apiKey, err := parseAPIKey()
	if err != nil {
		return err
	}

	tmdbClient, err := newTMDBClient(jackett.NewJackett(&jackett.Settings{
		ApiURL: jackettURL,
		ApiKey: apiKey,
	}), logger)
	if err != nil {
		return err
	}

	for {
		err = tmdbClient.Discover()
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}