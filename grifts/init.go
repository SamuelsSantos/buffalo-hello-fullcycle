package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/learn-jorney/hello_fullcycle/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
