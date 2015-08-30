package auth

import (
	"strings"
	"time"

	"github.com/robvdl/gcms/config"
	"github.com/robvdl/gcms/db"
)

// User is a user that can log into the cms
type User struct {
	db.Model
	Username    string `sql:"unique_index"`
	FirstName   string
	LastName    string
	Email       string
	Password    string
	IsActive    bool
	IsSuperuser bool
	LastLogin   time.Time
	Groups      []Group `gorm:"many2many:auth_user_group"`
}

// Group is a container for permissions
type Group struct {
	db.Model
	Name        string       `sql:"unique_index"`
	Permissions []Permission `gorm:"many2many:auth_group_permission"`
}

// Permission has a name and description
type Permission struct {
	db.Model
	Name        string `sql:"unique_index"`
	Description string `sql:"type:text"`
}

// TableName returns the table name gorm should use for the User model
func (u *User) TableName() string {
	return "auth_user"
}

// TableName returns the table name gorm should use for the Group model
func (g *Group) TableName() string {
	return "auth_group"
}

// TableName returns the table name gorm should use for the Permission model
func (p *Permission) TableName() string {
	return "auth_permission"
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
