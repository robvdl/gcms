package admin

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func AdminPage(c *gin.Context) {
	c.HTML(200, "admin.html", pongo2.Context{
		"title": "Admin",
	})
}
