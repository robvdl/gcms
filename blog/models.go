package blog

import "github.com/robvdl/gcms/db"

// Blog has a name and description and holds a number of Post objects.
type Blog struct {
	db.Model
	Name         string `sql:"unique_index"`
	Destcription string `sql:"type:text"`
	Posts        []Post `gorm:"many2many:blog_blog_post"`
}

// Post is a blog post, it can be used in multiple Blogs.
type Post struct {
	db.Model
	Title      string
	Slug       string     `sql:"unique_index"`
	Body       string     `sql:"type:text"`
	Categories []Category `gorm:"many2many:blog_post_category;"`
}

// Category is used to group Post objects.
type Category struct {
	db.Model
	Name        string `sql:"unique_index"`
	Slug        string `sql:"unique_index"`
	Description string `sql:"type:text"`
}

// TableName returns the table name gorm should use for the Blog model.
func (b *Blog) TableName() string {
	return "blog_blog"
}

// TableName returns the table name gorm should use for the Post model.
func (p *Post) TableName() string {
	return "blog_post"
}

// TableName returns the table name gorm should use for the Category model.
func (c *Category) TableName() string {
	return "blog_category"
}
