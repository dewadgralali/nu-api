package repository

import (
	"fmt"
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

// Push stores the data.
func (repo *CategoryRepository) Push(data *model.Category) error {
	return repo.db.Create(data).Error
}

// FindBy returns category by field-value.
func (repo *CategoryRepository) FindBy(field string, value interface{}) model.Category {
	category := model.Category{}

	repo.db.Where(fmt.Sprintf("%s = ?", field), value).Find(&category)

	return category
}

// Find returns category by ID.
func (repo *CategoryRepository) Find(id uint) model.Category {
	return repo.FindBy("id", id)
}

// Update updates category model.
func (repo *CategoryRepository) Update(data *model.Category) error {
	return repo.db.Save(data).Error
}
