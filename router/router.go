package router

import "github.com/gin-gonic/gin"

// NewRouter creates a default gin mux and sets the routes required.
func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return router
}
