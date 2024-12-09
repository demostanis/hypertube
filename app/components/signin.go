package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Signin() Node {
	return Form(Method("Post"), Action("/signin"),
		Center(
			P(Class("title"), Text("Sign in to Crocotube")),
			Input(Class("input"),
				Placeholder("Username..."),
				Name("username")),

			Div(Class("columns"),
				Div(Class("column"),
					Input(Class("input"),
						Placeholder("First Name..."),
						Name("firstName")),
				),
				Div(Class("column"),
					Input(Class("input"),
						Placeholder("Last Name..."),
						Name("lastName")),
				),
			),

			Input(Class("input"),
				Type("email"),
				Placeholder("Email..."),
				Name("email")),
			Input(Class("input password"),
				Type("password"),
				Placeholder("Password..."),
				Name("password")),
			Input(Class("input password"),
				Type("password"),
				Placeholder("Confirm Password..."),
				Name("passwordCheck")),
			Button(Class("button"),
				Text("Sign in"),
			),
		),
	)
}
