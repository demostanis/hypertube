package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

var bulma = "https://cdn.jsdelivr.net/npm/bulma@1.0.2/css/bulma.min.css"

func PageWithTitle(title string, content ...Node) Node {
	return HTML5(HTML5Props{
		Title: title,
		Head: []Node{
			Link(
				Rel("stylesheet"),
				Href(bulma),
			),
			Link(
				Rel("stylesheet"),
				Href("/static/css/styles.css"),
			),
		},
		Body: content,
	})
}

func Page(content ...Node) Node {
	return PageWithTitle("Crocotube", content...)
}
