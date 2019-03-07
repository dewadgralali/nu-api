package repository

import "github.com/jinzhu/gorm"

// NewCategoryRepository returns new CategoryRepository instance
func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}
