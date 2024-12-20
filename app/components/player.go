package components

import (
	"github.com/demostanis/hypertube/mvdb"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"strconv"
)

func getCardImage(imagePath string, id int, isMovie bool) Node {
	var trailerLink string
	var cardImage Node

	if id != -1 {
		trailerLink = mvdb.GetTrailer(strconv.Itoa(id), isMovie)
	} else {
		trailerLink = ""
	}

	if trailerLink == "" {
		cardImage = Figure(
			Class("image is-9by10"),
			Img(
				Attr("style", "width: 100%; aspect-ratio: 16 / 9;"),
				Src("https://image.tmdb.org/t/p/original/"+imagePath),
			),
		)
	} else {
		cardImage = IFrame(
			Attr("style", "width: 100%; aspect-ratio: 16 / 9;"),
			Src(trailerLink),
			Attr("allowfullscreen"),
		)
	}

	return cardImage
}

func getCardButton(ep, isMovie bool) Node {
	var button Node

	if ep || isMovie {
		button = Button(
			Class("card-button"),
			Span(
				Class("icon"), Attr("style", "color: mediumslateblue;"),
				I(
					Class("fa-solid fa-play fa-2xl"),
					Attr("aria-hidden", "true"),
				),
			),
		)
	} else {
		button = nil
	}

	return button
}

func Player(name, imagePath, overview string, id int, isMovie, ep bool) Node {
	if isMovie {
		id = -1
	}
	trailer := getCardImage(imagePath, id, isMovie)
	button := getCardButton(ep, isMovie)

	return Div(Class("card player"),
		Attr("style", "height: 100%;"),
		Header(
			Attr("style", "display: block; text-align: center;"),
			Class("card-header-title"),
			Text(name),
		),
		button,
		trailer,
		Div(
			Class("card-content"),
			Text(overview),
			Attr("style",
				"background-color: var(--bulma-card-background-color);"),
		),
	)
}
