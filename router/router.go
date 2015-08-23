package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/robvdl/gcms/config"
	c "github.com/robvdl/gcms/controllers"
	m "github.com/robvdl/gcms/middleware"
)

// NewRouter creates a default gin mux and sets the routes required.
// Here we setup the application middleware and add the application
// routes for the web application, then return the gin mux object.
func NewRouter() *gin.Engine {
	r := gin.Default()
	setupMiddleware(r)
	setupRoutes(r)
	return r
}

// setupMiddleware is an internal method where we setup middleware
func setupMiddleware(r *gin.Engine) {
	// TODO: CACHE_URL should come from an environment variable but this requires
	// validating and parsing of the connection url into it's base components.
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte(config.Config.Session_Secret))
	r.Use(sessions.Sessions("session", store))

	r.Use(m.AuthMiddleware())
}

// setupRoutes is the internal method where we add our routes
func setupRoutes(r *gin.Engine) {
	r.GET("/", c.APIStatus)
	r.GET("/login", c.APILogin)
	r.GET("/logout", c.APILogout)
}
