package cmd

import (
	"github.com/codegangsta/cli"

	"github.com/robvdl/gcms/db"
)

// CmdMigrate runs gorm AutoMigrate to create the database tables
var CmdMigrate = cli.Command{
	Name:        "migrate",
	Usage:       "Runs gorm automigrate to create database tables.",
	Description: "Run gorm automigrate on the models to create database tables.",
	Action:      migrate,
	Flags:       []cli.Flag{},
}

func migrate(ctx *cli.Context) {
	db.Migrate()
}
