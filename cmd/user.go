package cmd

import (
	"github.com/codegangsta/cli"

	"github.com/robvdl/gcms/auth"
	"github.com/robvdl/gcms/db"
)

// CmdWeb starts the web server
var CmdCreateSuperuser = cli.Command{
	Name:        "createsuperuser",
	Usage:       "Create a new superuser",
	Description: "Creates a new superuser who can log into admin.",
	Action:      createSuperuser,
	Flags:       []cli.Flag{},
}

// create a hardcoded superuser for now
func createSuperuser(ctx *cli.Context) {
	user := auth.User{Username: "admin", IsActive: true, IsSuperuser: true}
	user.SetPassword("password")
	db.DB.Create(&user)
}
