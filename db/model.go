package db

import "time"

// Model is the base struct for all models using struct embedding.
// Gorm also provides it's own Model base struct but this is not used,
// because I am not interested in using the soft deletes feature of Gorm.
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
