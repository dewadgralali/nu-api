package repository

import (
	"fmt"
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

// Push stores the data.
func (repo *PostRepository) Push(data *model.Post) error {
	return repo.db.Create(data).Error
}

// FindBy returns post by field-value.
func (repo *PostRepository) FindBy(field string, value interface{}) model.Post {
	post := model.Post{}

	repo.db.Where(fmt.Sprintf("%s = ?", field), value).First(&post)

	return post
}

// Find returns post by ID.
func (repo *PostRepository) Find(id uint) model.Post {
	return repo.FindBy("id", id)
}

// Update updates post model.
func (repo *PostRepository) Update(data *model.Post) error {
	return repo.db.Save(data).Error
}

// Delete deletes post by ID.
func (repo *PostRepository) Delete(id uint) error {
	return repo.db.Where("id = ?", id).Delete(&model.Post{}).Error
}
