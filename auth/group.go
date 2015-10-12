package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/manyminds/api2go/jsonapi"
)

// Group is a container for permissions
type Group struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        string       `sql:"unique_index"`
	Permissions []Permission `gorm:"many2many:auth_group_permission" jsonapi:"-"`
}

// TableName returns the table name gorm should use for the Group model
func (g *Group) TableName() string {
	return "auth_group"
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
