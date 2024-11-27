package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Contents(content Node) Node {
	return Section(Class("hero is-fullheight-with-navbar"),
		Div(Class("hero-body"),
			Div(Class("container"),
				content)))
}
