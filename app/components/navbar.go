package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func navbarItemStart(name string, path string) Node {
	return A(Class("navbar-item"), Attr("style", "border-right: solid 1px mediumslateblue; color: mediumslateblue;"), Href(path), Text(name))
}

func navbarItemEnd(name string, path string) Node {
	return A(Class("navbar-item"), Attr("style", "border-left: solid 1px mediumslateblue; color: mediumslateblue;"), Href(path), Text(name))
}

func Navbar() Node {
	return Nav(Class("navbar"),
		Div(Class("navbar-menu"),
			Div(Class("navbar-start"),
				navbarItemStart("Home", "/"),
				navbarItemStart("Videos", "/videos"),
			),
			Div(Class("navbar-end"),
				navbarItem("Login", "/login"),
				navbarItem("Signin", "/signin"),
			),
		),
	)
}
