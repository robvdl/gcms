package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/robvdl/gcms/config"
)

// NewRouter creates a default gin mux and sets the routes required.
// It also sets up the application middlwares such as cache, we might
// want to consider refactoring this so it only sets up routes, and
// setting up application middlewares moved elsewhere.
func NewRouter() *gin.Engine {
	r := gin.Default()

	// TODO: CACHE_URL should come from an environment variable but this requires
	// validating and parsing of the connection url into it's base components.
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte(config.Config.Session_Secret))
	r.Use(sessions.Sessions("session", store))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return r
}
