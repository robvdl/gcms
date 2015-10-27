package auth

import (
	"strings"

	"github.com/robvdl/gcms/config"
	"github.com/robvdl/gcms/models"
)

// SetPassword creates a password hash and updates the user
func SetPassword(u *models.User, password string) {
	if config.Config.Password_Algorithm == "bcrypt" {
		u.Password = bcryptPasswordString(
			password,
			config.Config.Password_Cost,
		)
	} else if strings.HasPrefix(config.Config.Password_Algorithm, "pbkdf2") {
		u.Password = pbkdf2PasswordString(
			password,
			config.Config.Password_Algorithm,
			config.Config.Password_Iterations,
			pkbdf2GenSalt(config.Config.Password_Salt_Size),
		)
	}
}

// CheckPassword checks a password against the password hash stored
// on the user object.
func CheckPassword(u *models.User, password string) bool {
	if strings.HasPrefix(u.Password, "bcrypt") {
		return bcryptCheckPassword(u.Password, password)
	} else if strings.HasPrefix(u.Password, "pbkdf2") {
		return pbkdf2CheckPassword(u.Password, password)
	}
	return false
}
