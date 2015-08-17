package main

import (
	"os"
	"path"
	"runtime"

	"github.com/codegangsta/cli"
	_ "github.com/lib/pq"

	"github.com/robvdl/gcms/cmd"
	"github.com/robvdl/gcms/config"
	"github.com/robvdl/gcms/db"
)

// AppVersion is the application version, build this from Git tag later
const AppVersion = "0.1"

func init() {
	// As of Go 1.5 this will be the default so we won't need to do this anymore
	// Before Go 1.5, this actually defaults to 1 CPU unless you do this.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// load config file based in project name
	_, project := path.Split(os.Args[0])
	config.LoadAppConfig(project)

	// establish database connection after config file is loaded
	db.Connect()
	db.Migrate()

	// Defines a cli application
	app := cli.NewApp()
	app.Name = project
	app.Usage = "Content management system"
	app.Version = AppVersion
	app.Commands = []cli.Command{
		cmd.CmdWeb,
	}

	app.Run(os.Args)
}
