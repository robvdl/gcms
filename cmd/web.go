package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"

	"github.com/robvdl/gcms/config"
)

// CmdWeb starts the web server
var CmdWeb = cli.Command{
	Name:        "web",
	Usage:       "Start the web server",
	Description: "Run gcms web to start the server",
	Action:      runWeb,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Value:  "/etc/default/gcms",
			Usage:  "Configuration file path",
			EnvVar: "",
		},
	},
}

func runWeb(ctx *cli.Context) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run(":" + config.Config.Port)
}
