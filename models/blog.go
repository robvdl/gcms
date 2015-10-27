package models

import "time"

// Blog has a name and description and holds a number of BlogPost objects.
type Blog struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name         string `sql:"unique_index"`
	Destcription string `sql:"type:text"`
	Posts        []Post `gorm:"many2many:blog_post"`
}

// Post is a blog post, it can be used in multiple Blogs.
type Post struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Title      string
	Slug       string     `sql:"unique_index"`
	Body       string     `sql:"type:text"`
	Categories []Category `gorm:"many2many:post_category"`
}

// Category is used to group Post objects.
type Category struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        string `sql:"unique_index"`
	Slug        string `sql:"unique_index"`
	Description string `sql:"type:text"`
}
