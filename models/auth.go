package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
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
	hashAlg := "pbkdf2-sha256"
	salt := make([]byte, 10)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// TODO: allow switching between pbkdf2-sha256 and bcrypt eventually
	iterations := 12000
	hash := pbkdf2.Key([]byte(password), salt, iterations, sha256.Size, sha256.New)
	u.Password = createPasswordString(hashAlg, iterations, salt, hash)
}

// CheckPassword checks a password against the password hash stored
// on the user object, for the time being it can only do pbkdf2-sha256
func (u *User) CheckPassword(password string) bool {
	parts := strings.Split(u.Password, "$")
	hashAlg := parts[0]

	iterations, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	salt, err := hex.DecodeString(parts[2])
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	hash, err := hex.DecodeString(parts[3])
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	hash = pbkdf2.Key([]byte(password), salt, iterations, sha256.Size, sha256.New)
	return u.Password == createPasswordString(hashAlg, iterations, salt, hash)
}

func createPasswordString(hashAlg string, iterations int, salt, hash []byte) string {
	return fmt.Sprintf("%s$%d$%x$%x", hashAlg, iterations, salt, hash)
}
