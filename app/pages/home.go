package pages

import (
	"github.com/demostanis/hypertube/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Home() Node {
	return components.Page(
		Class("has-navbar-fixed-top"),
		components.Navbar(),
		components.HeadLine("https://api.themoviedb.org/3/movie/top_rated?language=fr-FR&page=1"),
		components.FilmCategory("https://api.themoviedb.org/3/movie/popular?language=fr-FR&page=1&region=fr-FR", "Popular Movies"),
		components.FilmCategory("https://api.themoviedb.org/3/tv/popular?language=fr-FR&page=1&region=fr-FR", "Popular Series"),
		Div(ID("film-popup")),
		components.Foot())
}
