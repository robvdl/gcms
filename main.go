package main

import (
	"os"
	"runtime"

	"github.com/codegangsta/cli"
	_ "github.com/lib/pq"

	"github.com/robvdl/gcms/cmd"
	"github.com/robvdl/gcms/config"
	"github.com/robvdl/gcms/db"
)

func init() {
	// As of Go 1.5 this will be the default so we won't need to do this anymore
	// Before Go 1.5, this actually defaults to 1 CPU unless you do this.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Connect to the database, the connection is stored in db.DB
	// This allows the connection to be imported and used by sub-modules.
	db.Connect()

	// Migrate runs gorms AutoMigrate function on all the models
	db.Migrate()

	// Defines a cli application
	app := cli.NewApp()
	app.Name = config.AppName
	app.Usage = "Content management system"
	app.Version = config.AppVersion
	app.Commands = []cli.Command{
		cmd.CmdWeb,
	}

	app.Run(os.Args)
}
