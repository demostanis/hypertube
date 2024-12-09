package components

import (
	"encoding/json"
	"fmt"
	"io"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"net/http"
	"os"
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

// func CreateCardGrill(FilmList ApiResponse) Node {
// 	// Create an array large enough to hold all movies
// cards := make([]Node, len(FilmList.Results))

// // Fill the array with all movies
// for i, movie := range FilmList.Results {
// 	cards[i] = Div(Class("column"), Div(
// 		Class("cell"),
// 		Card(movie.OriginalTitle, movie.PosterPath),
// 	),)
// }

// 	return Div(Class("category"),
// 	Div(Class("control arrows"),
// 		Button(Class("arrow-left"), Text("◀"),
// 			Attr("onclick", "scrollGridLeft()"),
// 		),
// 		Button(Class("arrow-right"), Text("▶"),
// 			Attr("onclick", "scrollGridRight()"),
// 		),
// 	),
// 	Div(Class("columns is-mobile"),
// 		Div(append([]Node{Class("grid is-column-gap-4.5"), Attr("style", "overflow-x: auto; flex-wrap: nowrap; margin: 0;")}, cards[:]...)...),
// 	),
// 	Script(Src("/static/js/scroll.js")),
// )
// }

func CreateCardGrill(FilmList ApiResponse) Node {
	cards := make([]Node, len(FilmList.Results))

	// Fill the array with all movies
	for i, movie := range FilmList.Results {
		cards[i] = Div(Class("column"), Div(
			Class("cell"),
			Card(movie.OriginalTitle, movie.PosterPath),
		))
	}

	return Div(Class("category"),
		Div(Class("control arrows"),
			Button(Class("arrow-left"), Text("◀"), Attr("onclick", "scrollGridLeft()")),
			Button(Class("arrow-right"), Text("▶"), Attr("onclick", "scrollGridRight()")),
		),
		Div(
			append([]Node{Class("columns is-mobile"), Attr("style", "overflow-x: auto; flex-wrap: nowrap; margin: 0;")}, cards[:]...)...,
		),
		Script(Src("/static/js/scroll.js")),
	)
}

func CardGrill() Node {
	var PopularList ApiResponse

	err := json.Unmarshal([]byte(GetPopularMovies()), &PopularList)
	if err != nil {
		fmt.Println("Erreur lors du parsing JSON :", err)
		return Div(Text("Erreur lors de la récupération des films."))
	}
	return CreateCardGrill(PopularList)
}
