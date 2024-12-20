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

func SearchTrailer(ContentID, Lang string) string {
	var Trailers ApiResponse

	response := CallMvdbDefault("https://api.themoviedb.org/3/movie/" + ContentID + "/videos" + Lang)

	err := json.Unmarshal(response, &Trailers)
	if err != nil {
		return ""
	}

	for _, Trailer := range Trailers.Results {
		if Trailer.TrailerOfficial &&
			Trailer.TrailerSite == "YouTube" &&
			Trailer.TrailerSize == 1080 &&
			(Trailer.TrailerType == "Teaser" || Trailer.TrailerType == "Trailer") {
			return "https://www.youtube.com/embed/" + Trailer.TrailerKey
		}
	}
	return ""
}

func FindTrailer(ContentID string) string {
	TrailerLink := SearchTrailer(ContentID, "?language=fr-FR")
	if TrailerLink == "" {
		return SearchTrailer(ContentID, "")
	}
	return TrailerLink
}
