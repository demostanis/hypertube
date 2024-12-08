package mvdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Movie struct {
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
	Results []Movie `json:"results"`
}

func CallMvdbDefault(link string) string {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête :", err)
		return ""
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("MOVIE_DB_API"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête :", err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse :", err)
		return ""
	}

	return string(body)
}

func SearchTrailer(FilmID, Lang string) string {
	var Trailers ApiResponse

	response := CallMvdbDefault("https://api.themoviedb.org/3/movie/" + FilmID + "/videos" + Lang)

	json.Unmarshal([]byte(response), &Trailers)

	for _, Trailer := range Trailers.Results {
		if Trailer.TrailerOfficial && Trailer.TrailerSite == "YouTube" && Trailer.TrailerSize == 1080 && (Trailer.TrailerType == "Teaser" || Trailer.TrailerType == "Trailer") {
			return "https://www.youtube.com/embed/" + Trailer.TrailerKey
		}
	}
	return ""
}

func FindTrailer(FilmID string) string {
	TrailerLink := SearchTrailer(FilmID, "?language=fr-FR")
	if TrailerLink == "" {
		return SearchTrailer(FilmID, "")
	}
	return TrailerLink
}
