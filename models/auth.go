package models

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

// User is a user that can log into the cms
type User struct {
	Model
	Username    string `sql:"size:100;unique_index"`
	FirstName   string `sql:"size:100"`
	LastName    string `sql:"size:100"`
	Email       string `sql:"size:100"`
	Password    string `sql:"size:200"`
	IsActive    bool
	IsSuperuser bool
	LastLogin   time.Time
	Groups      []Group `gorm:"many2many:user_group"`
}

// Group is a container for permissions
type Group struct {
	Model
	Name        string       `sql:"size:100;unique_index"`
	Permissions []Permission `gorm:"many2many:group_permission"`
}

// Permission has a name and description
type Permission struct {
	Model
	Name        string `sql:"size:100;unique_index"`
	Description string `sql:"type:text"`
}

// SetPassword creates a password has and updates the user
// For the time being it is only capable of doing pbkdf2-sha256
func (u *User) SetPassword(password string) {
	salt := make([]byte, 10)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// TODO: allow switching between pbkdf2-sha256 and bcrypt eventually
	iterations := 12000
	hash := pbkdf2.Key([]byte(password), salt, iterations, sha256.Size, sha256.New)
	u.Password = fmt.Sprintf("$pbkdf2-sha256$%d$%x$%x", iterations, salt, hash)
}
