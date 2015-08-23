package auth

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/robvdl/gcms/config"
)

// bcryptPasswordString is an internal function that generates a bcrypt
// encoded password string.
func bcryptPasswordString(password string, cost int) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), config.Config.Password_Cost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return fmt.Sprintf("bcrypt%s", hashedPassword)
}
