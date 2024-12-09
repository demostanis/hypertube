package pages

import (
	"github.com/demostanis/hypertube/components"

	. "maragu.dev/gomponents"
)

func InternalError() Node {
	return components.Page(
		components.Navbar(),
		components.Contents(
			components.InternalError(),
		),
		components.Foot(),
	)
}
