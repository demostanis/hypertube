package pages

import (
	"github.com/demostanis/hypertube/components"

	. "maragu.dev/gomponents"
)

func Signin() Node {
	return components.Page(
		components.Navbar(),
		components.Contents(
			components.Signin(),
		),
		components.Foot(),
	)
}
