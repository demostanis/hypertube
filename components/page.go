package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

var bulma = "https://cdn.jsdelivr.net/npm/bulma@1.0.2/css/bulma.min.css"

func PageWithTitle(content Node, title string) Node {
	return HTML5(HTML5Props{
		Title: title,
		Head: []Node{
			Link(
				Rel("stylesheet"),
				Href(bulma),
			),
		},
		Body: []Node{content},
	})
}

func Page(content Node) Node {
	return PageWithTitle(content, "Hypertube")
}
