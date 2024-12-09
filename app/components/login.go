package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Center(content ...Node) Node {
	children := make([]Node, len(content))
	for _, n := range content {
		children = append(children, Div(Class("column"), n))
	}

	return Div(Class("columns is-1 is-centered has-text-centered"),
		Div(children...))
}

func Login() Node {
	return Form(Method("Post"), Action("/login"),
		Center(
			P(Class("title"), Text("Log in to Crocotube")),
			Input(Class("input"),
				Placeholder("Username..."),
				Name("username")),
			// TODO: there should be an eye icon inside this field
			Input(Class("input password"),
				Placeholder("Password..."),
				Name("password")),
			Button(Class("button"),
				Text("Log in"),
			),
		),
	)
}
