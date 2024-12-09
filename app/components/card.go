package components

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/demostanis/hypertube/mvdb"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var CategoryIndex = 0

var ScrollLeftAttr = `event.preventDefault();
		const container = document.getElementById('%s');
		const RightArrow = document.getElementById('%s-right');
		const LeftArrow = document.getElementById('%s-left');

		const opacity = (container.scrollLeft - container.offsetWidth <= 0) ? 0 : 1;
		LeftArrow.style.zIndex = opacity;
		RightArrow.style.zIndex = 1;
		LeftArrow.style.opacity = opacity;
		RightArrow.style.opacity = 1;
		
		container.scrollLeft -= container.offsetWidth;
		return false;`

var ScrollRightAttr = `event.preventDefault();
		const container = document.getElementById('%s');
		const RightArrow = document.getElementById('%s-right');
		const LeftArrow = document.getElementById('%s-left');
		const maxScrollLeft = container.scrollWidth - container.clientWidth;
		
		const opacity = (container.scrollLeft + container.offsetWidth >= maxScrollLeft) ? 0 : 1;
		RightArrow.style.zIndex = opacity;
		LeftArrow.style.zIndex = 1;
		RightArrow.style.opacity = opacity;
		LeftArrow.style.opacity = 1;
		
		container.scrollLeft += container.offsetWidth;
		return false;`

func Card(poster string) Node {
	return Div(Class("card"),
		Div(Class("card-image"),
			Figure(Class("image is-4by5"),
				Img(Class("poster-file"), Attr("style", "height: 121%;"), Src("https://image.tmdb.org/t/p/w500/"+poster)),
			),
		),
	)
}

func CreateCardGrill(FilmList mvdb.ApiResponse, categoryId string) Node {
	cards := make([]Node, len(FilmList.Results))

	for i, movie := range FilmList.Results {
		title := movie.Title
		title := movie.Title
		if title == "" {
			title = movie.Name
		}

		cards[i] = Div(Class("column pl-0 pr-5"), ID(categoryId+"-"+strconv.Itoa(i)),
			Attr("style", "display: flex;"),
			Div(
				Class("cell is-clickable"),
				Attr("hx-get", "/show-film-card"),
				Attr("hx-trigger", "click"),
				Attr("hx-target", "#film-card"),
				Attr("hx-swap", "innerHTML"),
				Attr("hx-vals", fmt.Sprintf(`{"filmId": %d, "titlefilm": "%s", "overview": "%s", "image": "%s"}`, movie.Id, title, movie.Overview, movie.ImagePath)),
				Card(movie.PosterPath),
			),
		)
	}

	return Div(Class("list"),
		Div(Class("control arrows is-hidden-mobile"),
			Button(
				Class("arrow-left"),
				ID(fmt.Sprintf("%s-left", categoryId)),
				Attr("onclick", fmt.Sprintf(ScrollLeftAttr, categoryId, categoryId, categoryId)),
				Span(Class("icon"),
					I(Class("fa-solid fa-chevron-left"), Attr("style", "scale: 1.5;")),
				),
			),
			Button(
				Class("arrow-right"),
				ID(fmt.Sprintf("%s-right", categoryId)),
				Attr("onclick", fmt.Sprintf(ScrollRightAttr, categoryId, categoryId, categoryId)),
				Span(Class("icon"),
					I(Class("fa-solid fa-chevron-right"), Attr("style", "scale: 1.5;")),
				),
			),
		),
		Div(
			append([]Node{
				Class("columns is-mobile pl-5"),
				ID(categoryId),
				Attr("style", "overflow-x: auto; flex-wrap: nowrap; margin: 0;scroll-behavior: smooth;position: relative;"),
			}, cards[:]...)...,
		),
	)
}

func CreateCategory(MovieList mvdb.ApiResponse, Name string) Node {
	categoryId := fmt.Sprintf("category-grid-%d", CategoryIndex)
	CategoryIndex = CategoryIndex + 1
	return Div(
		Class("category"),
		Div(
			Class("category-title title is-2 ml-5 mt-3"),
			Attr("style", "position: relative;"),
			Text(Name),
		),
		CreateCardGrill(MovieList, categoryId),
	)
}

func HeadLine(film mvdb.Movie) Node {
	Name := film.Title
	if Name == "" {
		Name = film.Name
	}
	return Div(
		Div(Class("headline is-hidden-mobile"),
			Attr("style", "height: 41vw;position: relative;"),
			Div(Class("headline-gradient-left")),
			Div(Class("headline-content"),
				Div(Class("headline-tilte"), Text(Name)),
				Div(Class("headline-overview"), Text(film.Overview)),
				Button(Class("button"), ID("play-button"),
					Span(Attr("style", "color: mediumslateblue;"), Text("PLAY")),
					Span(Class("icon"), Attr("style", "color: mediumslateblue;"),
						I(Class("fa-solid fa-play"), Attr("aria-hidden", "true")),
					),
				),
			),
			Div(Class("headline-gradient")),
			Img(Class("headline-img"), Src("https://image.tmdb.org/t/p/original"+film.ImagePath)),
		),
		Div(Class("headline-mobile is-hidden-tablet"),
			Div(Class("headline-mobile-gradient")),
			Div(Class("headline-mobile-content"),
				Div(Class("headline-mobile-tilte"), Text(Name)),
				Button(Class("button"), ID("play-mobile-button"),
					Span(Attr("style", "color: mediumslateblue;"), Text("PLAY")),
					Span(Class("icon"), Attr("style", "color: mediumslateblue;"),
						I(Class("fa-solid fa-play"), Attr("aria-hidden", "true")),
					),
				),
			),
			Img(Class("headline-mobile-img"), Src("https://image.tmdb.org/t/p/original"+film.PosterPath)),
		))
}

func CardGrill() Node {
	categories := []Node{}
	var TopRatedMovies mvdb.ApiResponse
	var PopularMovies mvdb.ApiResponse
	var PopularSeries mvdb.ApiResponse

	json.Unmarshal([]byte(mvdb.CallMvdbDefault("https://api.themoviedb.org/3/movie/top_rated?language=fr-FR&page=1")), &TopRatedMovies)
	json.Unmarshal([]byte(mvdb.CallMvdbDefault("https://api.themoviedb.org/3/movie/popular?language=fr-FR&page=1&region=fr-FR")), &PopularMovies)
	json.Unmarshal([]byte(mvdb.CallMvdbDefault("https://api.themoviedb.org/3/tv/popular?language=fr-FR&page=1&region=fr-FR")), &PopularSeries)

	categories = append(categories, HeadLine(TopRatedMovies.Results[0]))
	categories = append(categories, CreateCategory(PopularMovies, "Popular Movies"))
	categories = append(categories, CreateCategory(PopularSeries, "Popular Series"))
	categories = append(categories, Div(ID("film-card")))

	CategoryIndex = 0

	return Div(append([]Node{Class("categories-container")}, categories...)...)
}
