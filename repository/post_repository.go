package repository

import (
	"nu/model"

	"github.com/jinzhu/gorm"
)

// PostRepository represents repository for posts.
type PostRepository struct {
	db *gorm.DB
}

// Get returns all posts.
func (repo *PostRepository) Get() []model.Post {
	postList := make([]model.Post, 0)

	repo.db.Find(&postList)

	return postList
}
