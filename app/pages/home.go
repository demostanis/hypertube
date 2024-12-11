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
		components.ContentCategory("https://api.themoviedb.org/3/movie/popular?language=fr-FR&page=1&region=fr-FR", "Popular Movies"),
		components.ContentCategory("https://api.themoviedb.org/3/tv/top_rated?language=fr-FR&page=1", "Popular Series"),
		Div(ID("content-popup")),
		components.Foot())
}
