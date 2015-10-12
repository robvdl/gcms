package gallery

import "time"

// Album is a gallery album, it can only belong to one Gallery.
type Album struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	GalleryID   uint `sql:"index"`
	Name        string
	Slug        string  `sql:"unique_index"`
	Description string  `sql:"type:text"`
	Photos      []Photo `gorm:"many2many:gallery_album_photo;"`
}

// TableName returns the table name gorm should use for the Album model.
func (a *Album) TableName() string {
	return "gallery_album"
}
