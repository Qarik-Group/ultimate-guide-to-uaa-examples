package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	c.Set("email", c.Session().Get("email"))
	c.Set("scopes", c.Session().Get("scopes"))
	return c.Render(200, r.HTML("index.html"))
}
