package components

import (
	"fmt"
	"github.com/demostanis/hypertube/mvdb"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func epRuntime(runtime int) Node {
	hours := runtime / 60
	minutes := runtime % 60

	var formattedTime string
	if hours > 0 {
		formattedTime = fmt.Sprintf("%dh%dm", hours, minutes)
	} else {
		formattedTime = fmt.Sprintf("%dm", minutes)
	}

	return P(
		Class("runtime"),
		Text(formattedTime),
	)
}

func fillTab(
	episodes []mvdb.Episode,
	gridItems []Node,
	defaultImage string,
) []Node {
	var thumbnail string

	for i, episode := range episodes {
		if episode.Thumbnail != "" {
			thumbnail = episode.Thumbnail
		} else {
			thumbnail = defaultImage
		}
		gridItems[i] = Div(
			Class("cell"),
			Attr("hx-get", "/select-episode"),
			Attr("hx-trigger", "click"),
			Attr("hx-target", "#player-container"),
			Attr("hx-swap", "innerHTML"),
			Attr("hx-vals", fmt.Sprintf(
				`{"name": "E%d - %s", "overview": "%s", "thumbnail": "%s"}`,
				i+1, episode.Name, episode.Overview,
				"https://image.tmdb.org/t/p/original"+thumbnail)),
			Attr("style", "width: auto !important; position: relative;"),
			Button(
				Class("ep-button"),
				Attr("onclick", "window.location.href='#';"),
				Span(
					Class("icon"), Attr("style", "color: mediumslateblue;"),
					I(
						Class("fa-solid fa-play fa-lg"),
						Attr("aria-hidden", "true"),
					),
				),
			),
			Img(
				Src("https://image.tmdb.org/t/p/w500"+thumbnail),
				Attr("style", "aspect-ratio: 16/9; width: 100%;"),
			),
			epRuntime(episode.Runtime),
			Div(
				Attr("style", "z-index: 5"),
				Text(fmt.Sprintf("E%d - %s", i+1, episode.Name)),
			),
		)
	}

	return gridItems
}

func EpisodeGrid(episodes []mvdb.Episode, defaultImage string) Node {
	var gridItems []Node

	if len(episodes) < 6 {
		gridItems = make([]Node, 6)
	} else {
		gridItems = make([]Node, len(episodes))
	}

	gridItems = fillTab(episodes, gridItems, defaultImage)

	if len(episodes) < 6 {
		for i := len(episodes); i < 6; i++ {
			gridItems[i] = Div()
		}
	}

	epGrid := append(
		[]Node{
			Class("grid is-col-min-10"),
			Attr("style",
				"position: relative; padding: 6.4vw; padding-top: 0.7em;"),
		},
		gridItems...,
	)

	return Div(epGrid...)
}
