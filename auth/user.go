package auth

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go/jsonapi"

	"github.com/robvdl/gcms/config"
	"github.com/robvdl/gcms/db"
)

// User is a user that can log into the cms
// TODO: make Password field writeonly to the API
// TODO: make LastLogin field readonly to the API
type User struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Username  string `sql:"unique_index"`
	FirstName string
	LastName  string
	Email     string
	Password  string `jsonapi:"-"`
	Active    bool   `gorm:"column:is_active"`
	Superuser bool   `gorm:"column:is_superuser"`
	LastLogin time.Time
	Groups    []Group `gorm:"many2many:auth_user_group" jsonapi:"-"`
}

// TableName returns the table name gorm should use for the User model
func (u *User) TableName() string {
	return "auth_user"
}

// SetPassword creates a password has and updates the user
func (u *User) SetPassword(password string) {
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
	if strings.HasPrefix(u.Password, "bcrypt") {
		return bcryptCheckPassword(u.Password, password)
	} else if strings.HasPrefix(u.Password, "pbkdf2") {
		return pbkdf2CheckPassword(u.Password, password)
	}
	return false
}

// GetUserByUsername does a query that returns the User matching that username.
func GetUserByUsername(username string) *User {
	var user User
	db.DB.Where("username = ?", username).First(&user)
	return &user
}

// AuthenticatedUser returns the logged in user object or nil if not logged in.
func AuthenticatedUser(c *gin.Context) *User {
	// user comes from UserMiddleware
	v := c.MustGet("user")
	if v == nil {
		return nil
	}
	return v.(*User)
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u User) GetID() string {
	return strconv.FormatUint(uint64(u.ID), 10)
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *User) SetID(id string) error {
	value, err := strconv.ParseUint(id, 10, 32)
	if err == nil {
		u.ID = uint(value)
	}
	return err
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u User) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "groups",
			Name: "groups",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u User) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	for _, group := range u.Groups {
		result = append(result, jsonapi.ReferenceID{
			ID:   strconv.FormatUint(uint64(group.ID), 10),
			Type: "groups",
			Name: "groups",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u User) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}
	for key := range u.Groups {
		result = append(result, u.Groups[key])
	}

	return result
}

// SetToManyReferenceIDs sets the groups reference IDs and satisfies the
// jsonapi.UnmarshalToManyRelations interface
func (u *User) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "groups" {
		// TODO
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs adds new groups to users
func (u *User) AddToManyIDs(name string, IDs []string) error {
	if name == "groups" {
		// TODO
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// DeleteToManyIDs removes some groups form users
func (u *User) DeleteToManyIDs(name string, IDs []string) error {
	if name == "groups" {
		// TODO
	}

	return errors.New("There is no to-many relationship with the name " + name)
}
