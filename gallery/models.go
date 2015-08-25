package gallery

import "github.com/robvdl/gcms/db"

// Gallery holds a number of Album objects.
type Gallery struct {
	db.Model
	Name         string  `sql:"unique_index"`
	Destcription string  `sql:"type:text"`
	Albums       []Album // One-To-Many relationship to Album
}

// Album is a gallery album, it can only belong to one Gallery.
type Album struct {
	db.Model
	GalleryID   uint `sql:"index"`
	Name        string
	Slug        string  `sql:"unique_index"`
	Description string  `sql:"type:text"`
	Photos      []Photo // One-To-Many relationship to Album
}

// Photo is a photo object in an Album, it can only be in one Album
type Photo struct {
	db.Model
	AlbumID     uint `sql:"index"`
	Filename    string
	Description string `sql:"type:text"`
}

// TableName returns the table name gorm should use for the Album model
func (a *Album) TableName() string {
	return "gallery_album"
}

// TableName returns the table name gorm should use for the Photo model
func (p *Photo) TableName() string {
	return "gallery_photo"
}
