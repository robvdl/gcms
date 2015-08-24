package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/robvdl/gcms/auth"
	"github.com/robvdl/gcms/config"
)

// CmdWeb starts the web server
var CmdWeb = cli.Command{
	Name:        "web",
	Usage:       "Start the web server",
	Description: "Run gcms web to start the server",
	Action:      runWeb,
	Flags:       []cli.Flag{},
}

// setupMiddleware is an internal method where we setup GIN middleware
func setupMiddleware(r *gin.Engine) {
	// TODO: CACHE_URL should come from an environment variable but this requires
	// validating and parsing of the connection url into it's base components.
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte(config.Config.Session_Secret))
	r.Use(sessions.Sessions("session", store))
	r.Use(auth.UserMiddleware())
}

// setupRoutes is an internal method where we setup application routes
func setupRoutes(r *gin.Engine) {
	// TODO: we probably need something temporary for the / URL

	// session is a special api resource with POST and DELETE endpoints
	session := r.Group("/api/session")
	{
		session.POST("", auth.Login)
		session.DELETE("", auth.Logout)
	}
}

// runWeb is an starts the GIN application
func runWeb(ctx *cli.Context) {
	r := gin.Default()
	setupMiddleware(r)
	setupRoutes(r)
	r.Run(":" + config.Config.Port)
}
