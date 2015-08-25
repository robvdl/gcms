package blog

import "github.com/robvdl/gcms/db"

// Blog has a name and description and holds a number of Post objects.
type Blog struct {
	db.Model
	Name         string `sql:"unique_index"`
	Destcription string `sql:"type:text"`
	Posts        []Post // One-To-Many relationship to Post
}

// Category is used to group Post objects
type Category struct {
	db.Model
	Name        string `sql:"unique_index"`
	Slug        string `sql:"unique_index"`
	Description string `sql:"type:text"`
}

// Post is a blog post, it can only belong to one Blog using one to many.
type Post struct {
	db.Model
	BlogID     uint `sql:"index"`
	Title      string
	Slug       string     `sql:"unique_index"`
	Body       string     `sql:"type:text"`
	Categories []Category `gorm:"many2many:blog_post_category;"`
}

// TableName returns the table name gorm should use for the Post model
func (p *Post) TableName() string {
	return "blog_post"
}

// TableName returns the table name gorm should use for the Category model
func (c *Category) TableName() string {
	return "blog_category"
}
