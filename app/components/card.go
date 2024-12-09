package components

import (
	"encoding/json"
	"fmt"
	"io"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"net/http"
)

type Movie struct {
	PosterPath    string `json:"poster_path"`
	OriginalTitle string `json:"original_title"`
}

type ApiResponse struct {
	Results []Movie `json:"results"`
}

func GetPopularMovies() string {
	url := "https://api.themoviedb.org/3/movie/popular?language=fr&page=1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête :", err)
		return ""
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI5YWE4M2U4NzgwNTIzZmU2MDk2YjFlNTAwNGJmZDg0MyIsIm5iZiI6MTczMjc0NTAyNS41MTQxMTMyLCJzdWIiOiI2NzQ3OTY3MjY3OGExYmIzYmU0ZmVhNjIiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.l_-eMD4uXqxVPT5HJOQCLhnVo-JkemGGCK3JKlz3Ync")

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

func Card(name, poster string) Node {
	return Div(Class("card"),
		Div(Class("card-image"),
			Figure(Class("image is-4by5"),
				Img(Src("https://image.tmdb.org/t/p/w500/"+poster)),
			),
		),
		Div(Class("card-content"),
			Text(name),
		),
	)
}

var PopularMoviesDisplayed = 0

func CreateCardGrill() Node {
	var response ApiResponse

	err := json.Unmarshal([]byte(GetPopularMovies()), &response)
	if err != nil {
		fmt.Println("Erreur lors du parsing JSON :", err)
		return Div(Text("Erreur lors de la récupération des films."))
	}

	var cards [7]Node
	for i, movie := range response.Results[PopularMoviesDisplayed:] {
		if i >= len(cards) {
			break
		}
		cards[i] = Div(Class("cell"), Card(movie.OriginalTitle, movie.PosterPath))
	}
	PopularMoviesDisplayed = PopularMoviesDisplayed + 7

	fmt.Println(cards[0])
	return Div(Class("fixed-grid has-7-cols mx-4"),
		P(Class("mt-6")),
		Div(append([]Node{Class("grid is-column-gap-4.5")}, cards[:]...)...),
	)
}

func CardGrill() Node {
	return CreateCardGrill()
}
