package models

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"log"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"

	"github.com/robvdl/gcms/config"
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

// list of supported password algorithms represented by a map[string]bool
var algorithms = map[string]bool{
	"bcrypt":        true,
	"pbkdf2-sha256": true,
	"pbkdf2-sha384": true,
	"pbkdf2-sha512": true,
}

// SetPassword creates a password has and updates the user
func (u *User) SetPassword(password string) {
	// There is no point continuing if the hash algorithm setting is wrong.
	if !algorithms[config.Config.Password_Algorithm] {
		log.Fatal("Unsupported password algorithm: " + config.Config.Password_Algorithm)
	}

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
func (u *User) CheckPassword(password string) bool {
	parts := strings.Split(u.Password, "$")
	hashAlg := parts[0]

	// if we don't know this algorithm, just return false
	if !algorithms[hashAlg] {
		return false
	}

	if hashAlg == "bcrypt" {
		hashedPassword := []byte(u.Password)[6:] // first 6 bytes are bcrypt
		return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) == nil
	} else if strings.HasPrefix(hashAlg, "pbkdf2") {
		iterations, err := strconv.Atoi(parts[1])
		if err != nil {
			return false
		}

		salt, err := hex.DecodeString(parts[2])
		if err != nil {
			return false
		}

		return u.Password == pbkdf2PasswordString(password, hashAlg, iterations, salt)
	}
	return false
}

// bcryptPasswordString is an internal function that generates a bcrypt
// encoded password string.
func bcryptPasswordString(password string, cost int) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), config.Config.Password_Cost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return fmt.Sprintf("bcrypt%s", hashedPassword)
}

// pbkdf2PasswordString is an internal function that creates an pkbdf2 encoded
// password string, it tries to follow the same encoded format as Django,
// the advantage is that you can easily import users from Django.
func pbkdf2PasswordString(password, hashAlg string, iterations int, salt []byte) string {
	var keyLength int
	var hashFunc func() hash.Hash

	if hashAlg == "pbkdf2-sha256" {
		keyLength = sha256.Size
		hashFunc = sha256.New
	} else if hashAlg == "pbkdf2-sha384" {
		keyLength = sha512.Size384
		hashFunc = sha512.New384
	} else if hashAlg == "pbkdf2-sha512" {
		keyLength = sha512.Size
		hashFunc = sha512.New
	} else {
		log.Fatal("Unsupported password algorithm: " + hashAlg)
	}

	key := pbkdf2.Key([]byte(password), salt, config.Config.Password_Iterations, keyLength, hashFunc)
	return fmt.Sprintf("%s$%d$%x$%x", hashAlg, iterations, salt, key)
}

// pbkdf2GenSalt is an internal function that generates a salt using
// crypto/rand of the given size, this is not needed for bcrypt.
func pkbdf2GenSalt(size int) []byte {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal(err.Error())
	}
	return salt
}
