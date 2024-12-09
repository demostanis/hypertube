package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func InternalError() Node {
	return Center(
		P(Class("title"), Text("Internal Server Error")),
	)
}
