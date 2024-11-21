package pages

import (
	"github.com/demostanis/hypertube/components"

	. "maragu.dev/gomponents"
)

func Home() Node {
	return components.Page(components.Navbar())
}
