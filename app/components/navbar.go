package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func navbarItemStart(name, path, icon string) Node {
	return A(
		Class("navbar-item"),
		Attr("style", "color: mediumslateblue; margin-right: 20px;"),
		Href(path),
		Span(
			Class("icon"),
			I(Class("fa-solid fa-"+icon)),
		),
		Text(name),
	)
}

func navbarItemEnd(name, path, icon string) Node {
	return A(
		Class("navbar-item"),
		Attr("style", "color: mediumslateblue; margin-left: 20px;"),
		Href(path),
		Span(
			Class("icon"),
			I(Class("fa-solid fa-"+icon)),
		),
		Text(name),
	)
}

func Navbar() Node {
	return Nav(Class("navbar is-fixed-top"),
		Attr("style", "background-color: #1f2226;"),
		Div(Class("navbar-menu"),
			Div(Class("navbar-start"),
				Attr("style", "margin-left: 68px"),
				navbarItemStart("Home", "/", "house"),
				navbarItemStart("Videos", "/videos", "film"),
			),
			Div(Class("navbar-end"),
				Attr("style", "margin-right: 68px"),
				navbarItemEnd("Resheach...", "/", "magnifying-glass"),
				navbarItemEnd("Login", "/login", "user"),
			),
		),
	)
}
