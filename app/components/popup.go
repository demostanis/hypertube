package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func CreatePopupHeader(FilmTitle string) Node {
	return Header(Class("modal-card-head"),
		P(
			Class("modal-card-title"),
			Attr("style", "text-align: center; padding-left: 20px;"),
			Text(FilmTitle),
		),
		Button(Class("delete"),
			Attr("hx-get", "/empty"),
			Attr("hx-target", "closest .modal"),
			Attr("hx-swap", "delete"),
			Attr("aria-label", "close"),
		),
	)
}

func CreatePopupTrailer(TrailerLink, FilmImage string) Node {
	Style := "background-color: rgba(20, 22, 26, 1); padding: 15px; padding-bottom: 0px;"

	if TrailerLink != "" {
		return IFrame(
			ID("trailer"),
			Src(TrailerLink),
			Attr("allowfullscreen"),
			Attr("frameborder", "0"),
			Attr("width", "640"),
			Attr("height", "360"),
			Attr("style", Style),
		)
	}
	if FilmImage != "" {
		return Img(
			Src("https://image.tmdb.org/t/p/original"+FilmImage),
			Attr("style", Style+"width: 640px; height: 360px;"),
		)
	}
	return Div(
		Text("No trailer was found about this content :("),
		Br(),
		Text("But don't worry! You can still watch it!"),
		Attr("style", Style+"width: 640px; height: 360; text-align:center"),
	)
}

func CreatePopupFooter() Node {
	return Footer(Class("modal-card-foot"),
		Div(Class("buttons"),
			Button(Class("button"), Attr("style", "outile: solid 1px mediumslateblue"),
				Span(Attr("style", "color: mediumslateblue;"), Text("PLAY")),
				Span(Class("icon"), Attr("style", "color: mediumslateblue;"),
					I(Class("fa-solid fa-play"), Attr("aria-hidden", "true")),
				),
			),
		),
	)
}

func CreatePopup(FilmTitle, FilmOverview, TrailerLink, FilmImage string) Node {
	return Div(Class("modal is-active"),
		Div(Class("modal-background")),
		Div(Class("modal-card"),
			CreatePopupHeader(FilmTitle),
			CreatePopupTrailer(TrailerLink, FilmImage),
			Section(Class("modal-card-body"),
				Text(FilmOverview),
			),
			CreatePopupFooter(),
		),
	)
}
