package auth

import (
	"strconv"
	"time"
)

// Permission has a name and description
type Permission struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        string `sql:"unique_index"`
	Description string `sql:"type:text"`
}

// TableName returns the table name gorm should use for the Permission model
func (p *Permission) TableName() string {
	return "auth_permission"
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
