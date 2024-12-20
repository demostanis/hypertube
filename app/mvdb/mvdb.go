package mvdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Content struct {
	ImagePath  string `json:"backdrop_path"`
	PosterPath string `json:"poster_path"`
	Title      string `json:"title"`
	Name       string `json:"name"`
	Overview   string `json:"overview"`
	Id         int    `json:"id"`

	TrailerType     string `json:"type"`
	TrailerSite     string `json:"site"`
	TrailerSize     int    `json:"size"`
	TrailerOfficial bool   `json:"official"`
	TrailerKey      string `json:"key"`
}

type ApiResponse struct {
	Results []Content `json:"results"`
}

func CallMvdbDefault(link string) []byte {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		fmt.Println("Error creating the request:", err)
		return nil
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("MOVIE_DB_API"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error executing the request:", err)
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return nil
	}

	return []byte(body)
}

func searchTrailer(contentId, lang string, isMovie bool) string {
	var trailers ApiResponse
	var response []byte

	if isMovie {
		response = CallMvdbDefault("https://api.themoviedb.org/3/movie/" + contentId + "/videos" + lang)
	} else {
		response = CallMvdbDefault("https://api.themoviedb.org/3/tv/" + contentId + "/videos" + lang)
	}

	json.Unmarshal(response, &trailers)

	for _, trailer := range trailers.Results {
		if trailer.TrailerOfficial &&
			trailer.TrailerSite == "YouTube" &&
			trailer.TrailerSize == 1080 &&
			(trailer.TrailerType == "Teaser" || trailer.TrailerType == "Trailer") {
			return "https://www.youtube.com/embed/" + trailer.TrailerKey
		}
	}
	return ""
}

func GetTrailer(contentId string, isMovie bool) string {
	trailerLink := searchTrailer(contentId, "?language=fr-FR", isMovie)
	if trailerLink == "" {
		return searchTrailer(contentId, "", isMovie)
	}
	return trailerLink
}

func GetContentByID(contentId, lang string, isMovie bool) Content {
	var content Content
	var response []byte

	if isMovie {
		response = CallMvdbDefault("https://api.themoviedb.org/3/movie/" + contentId + "?language=" + lang)
	} else {
		response = CallMvdbDefault("https://api.themoviedb.org/3/tv/" + contentId + "?language=" + lang)
	}

	json.Unmarshal(response, &content)
	if isMovie {
		content.Name = content.Title
	} else {
		content.Title = content.Name
	}

	return content
}
