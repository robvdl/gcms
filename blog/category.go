package blog

import "time"

// Category is used to group Post objects.
type Category struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        string `sql:"unique_index"`
	Slug        string `sql:"unique_index"`
	Description string `sql:"type:text"`
}

// TableName returns the table name gorm should use for the Category model.
func (c *Category) TableName() string {
	return "blog_category"
}
