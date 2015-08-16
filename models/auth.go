package models

import "time"

// User is a user that can log into the cms
type User struct {
	Model
	Username    string `sql:"size:50;unique_index"`
	FirstName   string `sql:"size:50"`
	LastName    string `sql:"size:50"`
	Email       string `sql:"size:100"`
	Password    string `sql:"size:100"`
	IsActive    bool
	IsSuperuser bool
	LastLogin   time.Time
	Groups      []Group `gorm:"many2many:user_group"`
}

// Group is a container for permissions
type Group struct {
	Model
	Name        string       `sql:"size:50;unique_index"`
	Permissions []Permission `gorm:"many2many:group_permission"`
}

// Permission has a name and description
type Permission struct {
	Model
	Name        string `sql:"size:50;unique_index"`
	Description string `sql:"type:text"`
}
