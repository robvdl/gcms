package models

import "time"

// Gallery holds a number of Album objects.
type Gallery struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name         string  `sql:"unique_index"`
	Destcription string  `sql:"type:text"`
	Albums       []Album // One-To-Many relationship to Album
}

// Album is a gallery album, it can only belong to one Gallery.
type Album struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	GalleryID   uint `sql:"index"`
	Name        string
	Slug        string  `sql:"unique_index"`
	Description string  `sql:"type:text"`
	Photos      []Photo `gorm:"many2many:album_photo"`
}

// Photo is a photo object in an Album, it can only be in one Album.
type Photo struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Filename    string
	Description string `sql:"type:text"`
}
