package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/howeyc/gopass"

	"github.com/robvdl/gcms/auth"
	"github.com/robvdl/gcms/db"
)

// CmdCreateUser creates a new superuser
var CmdCreateUser = cli.Command{
	Name:        "createuser",
	Usage:       "Create a new superuser",
	Description: "Creates a new superuser who can log into admin.",
	Action:      createUser,
	Flags:       []cli.Flag{},
}

// createUser creates a new superuser, asking for a username and password.
func createUser(ctx *cli.Context) {
	var username, email, password, confirmPassword string

	for {
		fmt.Print("Username: ")
		_, err := fmt.Scanln(&username)
		if err != nil && err.Error() != "unexpected newline" {
			fmt.Println(err)
		}

		if username == "" {
			fmt.Println("Username cannot be blank")
		} else {
			break
		}
	}

	// email is optional
	fmt.Print("Email: ")
	_, err := fmt.Scanln(&email)
	if err != nil && err.Error() != "unexpected newline" {
		fmt.Println(err)
	}

	for {
		fmt.Print("Password: ")
		a := gopass.GetPasswd()
		fmt.Print("Confirm Password: ")
		b := gopass.GetPasswd()

		password = string(a)
		confirmPassword = string(b)

		if password == "" {
			fmt.Println("Password cannot be blank")
		} else if password != confirmPassword {
			fmt.Println("Passwords must match")
		} else {
			break
		}
	}

	user := auth.User{Username: username, Email: email, Active: true, Superuser: true}
	user.SetPassword(password)
	db.DB.Create(&user)
}
