package auth

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

	"golang.org/x/crypto/pbkdf2"

	"github.com/robvdl/gcms/config"
)

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

// pbkdf2CheckPassword checks a password hash against a password.
func pbkdf2CheckPassword(hashedPassword, password string) bool {
	parts := strings.Split(hashedPassword, "$")
	hashAlg := parts[0]

	iterations, err := strconv.Atoi(parts[1])
	if err != nil {
		return false
	}

	salt, err := hex.DecodeString(parts[2])
	if err != nil {
		return false
	}

	return hashedPassword == pbkdf2PasswordString(password, hashAlg, iterations, salt)
}
