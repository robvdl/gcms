package gallery

import "time"

// Photo is a photo object in an Album, it can only be in one Album.
type Photo struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Filename    string
	Description string `sql:"type:text"`
}

// TableName returns the table name gorm should use for the Photo model.
func (p *Photo) TableName() string {
	return "gallery_photo"
}
