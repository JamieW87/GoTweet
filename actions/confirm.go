package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func ConfirmHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("confirm.html"))
}
