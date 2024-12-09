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

func navbarItemMobil(pos, path, icon string) Node {
	return A(
		Attr("style", "position: absolute; color: mediumslateblue; top: 25%; left: "+pos+";"),
		Href(path),
		Span(
			Class("icon"),
			I(Class("fa-solid fa-"+icon)),
		),
	)
}

func Navbar() Node {
	return Nav(Class("navbar is-fixed-top"),
		Attr("style", "background-color: #1f2226;"),
		Div(Class("navbar-menu is-hidden-mobile"),
			Div(Class("navbar-start"),
				Attr("style", "margin-left: 68px"),
				navbarItemStart("Home", "/", "house"),
				navbarItemStart("Videos", "/videos", "film"),
			),
			Div(Class("navbar-end"),
				Attr("style", "margin-right: 68px"),
				navbarItemEnd("Research...", "/", "magnifying-glass"),
				navbarItemEnd("Login", "/login", "user"),
			),
		),
		Div(Class("navbar-menu is-hidden-tablet"),
			navbarItemMobil("15vw", "/", "house"),
			navbarItemMobil("35vw", "/videos", "film"),
			navbarItemMobil("60vw", "/", "magnifying-glass"),
			navbarItemMobil("80vw", "/login", "user"),
		),
	)
}
