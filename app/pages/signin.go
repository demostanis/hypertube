package pages

import (
	"github.com/demostanis/hypertube/components"

	. "maragu.dev/gomponents"
)

func Signin(err string) Node {
	return components.Page(
		components.Navbar(),
		components.Contents(
			components.Signin(err),
		),
		components.Foot(),
	)
}
