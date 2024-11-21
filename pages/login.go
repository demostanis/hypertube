package pages

import (
	"github.com/demostanis/hypertube/components"

	. "maragu.dev/gomponents"
)

func Login() Node {
	return components.Page(
		components.Navbar(),
		components.Contents(
			components.Login(),
		),
	)
}
