package components

import (
	"encoding/json"
	"github.com/demostanis/hypertube/mvdb"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func HeadlineDesktop(Film mvdb.Movie, Name string) Node {
	return Div(Class("headline is-hidden-mobile"),
		Attr("style", "height: 41vw;position: relative;"),
		Div(Class("headline-gradient-left")),
		Div(Class("headline-content"),
			Div(Class("headline-tilte"), Text(Name)),
			Div(Class("headline-overview"), Text(Film.Overview)),
			Button(Class("button"), ID("play-button"),
				Span(Attr("style", "color: mediumslateblue;"), Text("PLAY")),
				Span(Class("icon"), Attr("style", "color: mediumslateblue;"),
					I(Class("fa-solid fa-play"), Attr("aria-hidden", "true")),
				),
			),
		),
		Div(Class("headline-gradient")),
		Img(Class("headline-img"), Src("https://image.tmdb.org/t/p/original"+Film.ImagePath)),
	)
}

func HeadlineMobile(Film mvdb.Movie, Name string) Node {
	return Div(Class("headline-mobile is-hidden-tablet"),
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
		Img(Class("headline-mobile-img"), Src("https://image.tmdb.org/t/p/original"+Film.PosterPath)),
	)
}

func HeadLine() Node {
	var TopRatedMovies mvdb.ApiResponse

	json.Unmarshal([]byte(mvdb.CallMvdbDefault("https://api.themoviedb.org/3/movie/top_rated?language=fr-FR&page=1")), &TopRatedMovies)

	Name := TopRatedMovies.Results[0].Title
	if Name == "" {
		Name = TopRatedMovies.Results[0].Name
	}
	return Div(
		HeadlineDesktop(TopRatedMovies.Results[0], Name),
		HeadlineMobile(TopRatedMovies.Results[0], Name),
	)
}
