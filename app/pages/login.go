package pages

import (
	"github.com/demostanis/hypertube/components"

	. "maragu.dev/gomponents"
)

func Login(error string) Node {
	return components.Page(
		components.Navbar(),
		components.Contents(
			components.Login(error),
		),
		components.Foot(),
	)
}
