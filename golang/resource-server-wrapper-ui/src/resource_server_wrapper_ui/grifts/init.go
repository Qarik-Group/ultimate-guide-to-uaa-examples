package grifts

import (
	"resource_server_wrapper_ui/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
