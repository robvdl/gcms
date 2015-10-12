package blog

import "time"

// Blog has a name and description and holds a number of Post objects.
type Blog struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name         string `sql:"unique_index"`
	Destcription string `sql:"type:text"`
	Posts        []Post `gorm:"many2many:blog_blog_post"`
}

// TableName returns the table name gorm should use for the Blog model.
func (b *Blog) TableName() string {
	return "blog_blog"
}
