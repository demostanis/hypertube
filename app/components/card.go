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

var CategoryIndex = 0

type Movie struct {
	PosterPath    string `json:"poster_path"`
	OriginalTitle string `json:"original_title"`
	OriginalName  string `json:"original_name"`
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

func Card(name, poster string) Node {
	return Div(Class("card"),
		Div(Class("card-image"),
			Figure(Class("image is-4by5"),
				Img(Src("https://image.tmdb.org/t/p/w500/"+poster)),
			),
		),
		Div(Class("card-content"),
			Attr("style", "white-space: nowrap; overflow: hidden; text-overflow: ellipsis; text-align: center; color: white; font-weight: bold;"),
			Text(name),
		),
	)
}

func CreateCardGrill(FilmList ApiResponse, categoryId string) Node {
	cards := make([]Node, len(FilmList.Results))

	for i, movie := range FilmList.Results {
		title := movie.OriginalTitle
		if title == "" {
			title = movie.OriginalName
		}

		cards[i] = Div(Class("column pl-0 pr-5"), Div(
			Class("cell"),
			Card(title, movie.PosterPath),
		))
	}

	return Div(Class("list"),
		Div(Class("control arrows"),
			Button(Class("arrow-left"), Text("◀"),
				Attr("onclick", fmt.Sprintf("scrollGridLeft('%s')", categoryId))),
			Button(Class("arrow-right"), Text("▶"),
				Attr("onclick", fmt.Sprintf("scrollGridRight('%s')", categoryId))),
		),
		Div(
			append([]Node{
				Class("columns is-mobile pl-5"),
				ID(categoryId),
				Attr("style", "overflow-x: auto; flex-wrap: nowrap; margin: 0;"),
			}, cards[:]...)...,
		),
		Script(Src("/static/js/scroll.js")),
	)
}

func CreateCategory(MovieList ApiResponse, Name string) Node {
	categoryId := fmt.Sprintf("category-grid-%d", CategoryIndex)
	CategoryIndex = CategoryIndex + 1
	return Div(
		Class("category"),
		Div(
			Class("category-title title is-2 ml-5 mt-3"),
			Text(Name),
		),
		CreateCardGrill(MovieList, categoryId),
	)
}

func CardGrill() Node {
	categories := []Node{}
	var PopularMovies ApiResponse
	var PopularSeries ApiResponse

	err := json.Unmarshal([]byte(CallMvdbDefault("https://api.themoviedb.org/3/movie/popular?language=fr-FR&page=1")), &PopularMovies)
	if err != nil {
		fmt.Println("Erreur lors du parsing JSON :", err)
		return Div(Text("Erreur lors de la récupération des films."))
	}
	err1 := json.Unmarshal([]byte(CallMvdbDefault("https://api.themoviedb.org/3/tv/popular?language=fr-FR&page=1")), &PopularSeries)
	if err1 != nil {
		fmt.Println("Erreur lors du parsing JSON :", err1)
		return Div(Text("Erreur lors de la récupération des films."))
	}

	categories = append(categories, CreateCategory(PopularMovies, "Popular Movies"))
	categories = append(categories, CreateCategory(PopularSeries, "Popular Series"))

	return Div(append([]Node{Class("categories-container")}, categories...)...)
}
