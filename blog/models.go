package blog

import "github.com/robvdl/gcms/db"

// Blog has a name and description and holds a number of Post objects.
type Blog struct {
	db.Model
	Name         string
	Destcription string
	Posts        []Post // One-To-Many relationship to Post
}

// Category is used to group Post objects
type Category struct {
	db.Model
	Name        string
	Slug        string
	Description string
}

// Post is a blog post, it can only below to one Blog using one to many.
type Post struct {
	db.Model
	BlogID     uint `sql:"index"`
	Title      string
	Slug       string
	Body       string
	Categories []Category `gorm:"many2many:blog_post_category;"`
}

// TableName returns the table name gorm should use for the Post model
func (b *Post) TableName() string {
	return "blog_post"
}

// TableName returns the table name gorm should use for the Category model
func (c *Category) TableName() string {
	return "blog_category"
}
