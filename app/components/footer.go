package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Foot() Node {
	return Footer(Class("footer"),
		Div(Class("content has-text-centered"),
			Text("Crocotube by "),
			A(Href("https://github.com/demostanis"), Text("demostanis")),
			Text(", "),
			A(Href("https://github.com/DinoMalin"), Text("DinoMalin")),
			Text(" and "),
			A(Href("https://github.com/acasamit"), Text("acasamit")),
			Br(),
			A(Href("https://github.com/demostanis/hypertube"),
				Text("⭐ Star on GitHub")),
		),
	)
}
