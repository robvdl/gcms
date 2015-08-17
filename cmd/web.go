package cmd

import (
	"github.com/codegangsta/cli"

	"github.com/robvdl/gcms/config"
	"github.com/robvdl/gcms/router"
)

// CmdWeb starts the web server
var CmdWeb = cli.Command{
	Name:        "web",
	Usage:       "Start the web server",
	Description: "Run gcms web to start the server",
	Action:      runWeb,
	Flags:       []cli.Flag{},
}

func runWeb(ctx *cli.Context) {
	r := router.NewRouter()
	r.Run(":" + config.Config.Port)
}
