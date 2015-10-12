package gallery

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

// TableName returns the table name gorm should use for the Gallery model.
func (g *Gallery) TableName() string {
	return "gallery_gallery"
}
