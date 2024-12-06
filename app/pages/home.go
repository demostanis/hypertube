package pages

import (
	"github.com/demostanis/hypertube/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Home() Node {
	return components.Page(Class("has-navbar-fixed-top"), components.Navbar(), components.CardGrill(), components.Foot())
}
