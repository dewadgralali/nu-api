package model

import "time"

// Post represents a post
type Post struct {
	ID         uint
	Title      string
	Slug       string
	MDContent  string     `gorm:"type:text"`
	Categories []Category `gorm:"many2many:category_post"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
