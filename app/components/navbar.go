package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func navbarItem(name string, path string) Node {
	return A(Class("navbar-item"), Href(path), Text(name))
}

func Navbar() Node {
	return Nav(Class("navbar"),
		Div(Class("navbar-menu"),
			Div(Class("navbar-start"),
				navbarItem("Home", "/"),
				navbarItem("Videos", "/videos"),
			),
			Div(Class("navbar-end"),
				navbarItem("Login", "/login"),
			),
		),
	)
}
