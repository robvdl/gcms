package cmd

import (
	"github.com/codegangsta/cli"

	"github.com/robvdl/gcms/config"
	"github.com/robvdl/gcms/db"
	"github.com/robvdl/gcms/router"
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
			Value:  "/etc/default/" + config.AppName,
			Usage:  "Configuration file path",
			EnvVar: "",
		},
	},
}

func runWeb(ctx *cli.Context) {
	// loads the configuration file, can be overriden using a flag
	config.Load(ctx.String("config"))

	// establish database connection after config file is loaded
	db.Connect()
	db.Migrate()

	r := router.NewRouter()
	r.Run(":" + config.Config.Port)
}
