package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
	"github.com/robvdl/pongo2gin"

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

// csrfFailed is called by nosurf when the csrf token check fails
func csrfFailed(w http.ResponseWriter, r *http.Request) {
	// Set status code when overriding failure handler or it wil be 200
	w.WriteHeader(400)
	fmt.Fprintln(w, nosurf.Reason(r)) // reason of the failure
}

// runWeb is an starts the GIN application
func runWeb(ctx *cli.Context) {
	r := gin.Default()
	r.HTMLRender = pongo2gin.Default() // Use Pongo2 for templates

	setupMiddleware(r)
	setupRoutes(r)

	// Initialise nosurf for csrf token support.
	csrfHandler := nosurf.New(r)
	csrfHandler.SetFailureHandler(http.HandlerFunc(csrfFailed))
	csrfHandler.ExemptRegexp("/api/(.*)") // ignore API urls for the time being

	// Start the Gin application with nosurf (for csrf protection).
	// This is an alternative way to start up the Gin application.
	http.ListenAndServe(":"+config.Config.Port, csrfHandler)
}
