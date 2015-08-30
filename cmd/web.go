package cmd

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/robvdl/gcms/auth"
	"github.com/robvdl/gcms/config"
	"github.com/robvdl/pongo2gin"
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
	store, err := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte(config.Config.Session_Secret))
	if err != nil {
		log.Fatalln("Failed to connect to Redis.", err)
	}
	r.Use(sessions.Sessions("session", store))
	r.Use(auth.UserMiddleware())
}

// setupRoutes is an internal method where we setup application routes
func setupRoutes(r *gin.Engine) {
	// FIXME: this is just the temporary location of the login page
	r.GET("/", auth.LoginPage)
	r.POST("/", auth.LoginPage)

	// session is a special api resource with POST and DELETE endpoints
	session := r.Group("/api/session")
	{
		session.POST("", auth.LoginAPI)
		session.DELETE("", auth.LogoutAPI)
	}
}

// runWeb is an starts the GIN application
func runWeb(ctx *cli.Context) {
	r := gin.Default()
	r.HTMLRender = pongo2gin.Default() // Use Pongo2 for templates

	setupMiddleware(r)
	setupRoutes(r)

	r.Run(":" + config.Config.Port)
}
