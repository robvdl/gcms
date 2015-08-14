package router

import "github.com/gin-gonic/gin"

// NewRouter creates a default gin mux and sets the routes required.
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return r
}
