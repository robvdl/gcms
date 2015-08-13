package main

import (
	"os"
	"runtime"

	"github.com/codegangsta/cli"

	"github.com/robvdl/gcms/cmd"
)

// TODO: build this from Git tag instead
const AppVersion = "0.1"

func init() {
	// As of Go 1.5 this will be the default so we won't need to do this anymore
	// Before Go 1.5, this actually defaults to 1 CPU unless you do this.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "gcms"
	app.Usage = "content management system"
	app.Version = AppVersion
	app.Commands = []cli.Command{
		cmd.CmdWeb,
	}
	app.Run(os.Args)
}
