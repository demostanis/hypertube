package components

import (
	"encoding/json"

	"github.com/demostanis/hypertube/mvdb"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func HeadlineDesktop(Film mvdb.Film, Name string) Node {
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
		Img(
			Class("headline-img"),
			Src("https://image.tmdb.org/t/p/original"+Film.ImagePath),
		),
	)
}

func HeadlineMobile(Film mvdb.Film, Name string) Node {
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
		Img(
			Class("headline-mobile-img"),
			Src("https://image.tmdb.org/t/p/original"+Film.PosterPath),
		),
	)
}

func HeadLine(Request string) Node {
	var FilmList mvdb.ApiResponse

	json.Unmarshal(mvdb.CallMvdbDefault(Request), &FilmList)

	Name := FilmList.Results[0].Title
	if Name == "" {
		Name = FilmList.Results[0].Name
	}
	return Div(
		HeadlineDesktop(FilmList.Results[0], Name),
		HeadlineMobile(FilmList.Results[0], Name),
	)
}
