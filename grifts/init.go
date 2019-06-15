package grifts

import (
	"github.com/danoviedo91/todotasks/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
