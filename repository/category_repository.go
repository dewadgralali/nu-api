package repository

import (
	"nu/model"

	"github.com/jinzhu/gorm"
)

// CategoryRepository represents repository for categories.
type CategoryRepository struct {
	db *gorm.DB
}

// Get returns all categories.
func (repo *CategoryRepository) Get() []model.Category {
	categoryList := make([]model.Category, 0)

	repo.db.Find(&categoryList)

	return categoryList
}
