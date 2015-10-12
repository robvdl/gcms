package blog

import "time"

// Post is a blog post, it can be used in multiple Blogs.
type Post struct {
	ID        uint `gorm:"primary_key" jsonapi:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Title      string
	Slug       string     `sql:"unique_index"`
	Body       string     `sql:"type:text"`
	Categories []Category `gorm:"many2many:blog_post_category;"`
}

// TableName returns the table name gorm should use for the Post model.
func (p *Post) TableName() string {
	return "blog_post"
}
