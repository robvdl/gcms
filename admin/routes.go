package admin

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"

	"github.com/robvdl/gcms/auth"
)

// Admin is the cms admin page, it requires a login.
func Admin(c *gin.Context) {
	c.HTML(200, "admin.html", pongo2.Context{
		"title": "Admin",
		"user":  auth.AuthenticatedUser(c),
	})
}
