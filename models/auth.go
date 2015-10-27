package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/manyminds/api2go/jsonapi"
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
	Groups    []Group `gorm:"many2many:user_group" jsonapi:"-"`
}

// Group is a container for permissions
type Group struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        string       `sql:"unique_index"`
	Permissions []Permission `gorm:"many2many:group_permission" jsonapi:"-"`
}

// Permission has a name and description
type Permission struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        string `sql:"unique_index"`
	Description string `sql:"type:text"`
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

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (g Group) GetID() string {
	return strconv.FormatUint(uint64(g.ID), 10)
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (g *Group) SetID(id string) error {
	value, err := strconv.ParseUint(id, 10, 32)
	if err == nil {
		g.ID = uint(value)
	}
	return err
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (g Group) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "permissions",
			Name: "permissions",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (g Group) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	for _, permission := range g.Permissions {
		result = append(result, jsonapi.ReferenceID{
			ID:   strconv.FormatUint(uint64(permission.ID), 10),
			Type: "permissions",
			Name: "permissions",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (g Group) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}
	for key := range g.Permissions {
		result = append(result, g.Permissions[key])
	}

	return result
}

// SetToManyReferenceIDs sets the permissions reference IDs and satisfies the
// jsonapi.UnmarshalToManyRelations interface
func (g *Group) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "permissions" {
		// TODO
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs adds new permissions to groups
func (g *Group) AddToManyIDs(name string, IDs []string) error {
	if name == "permissions" {
		// TODO
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// DeleteToManyIDs removes some permissions form groups
func (g *Group) DeleteToManyIDs(name string, IDs []string) error {
	if name == "permissions" {
		// TODO
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (p Permission) GetID() string {
	return strconv.FormatUint(uint64(p.ID), 10)
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (p *Permission) SetID(id string) error {
	value, err := strconv.ParseUint(id, 10, 32)
	if err == nil {
		p.ID = uint(value)
	}
	return err
}
