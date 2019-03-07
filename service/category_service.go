package service

import (
	"fmt"
	"nu/model"
	"nu/repository"
)

// CategoryService represents service layer for
// category
type CategoryService struct {
	repo repository.CategoryRepositoryContract
}

// Get returns available categories
func (srv *CategoryService) Get() []model.Category {
	return srv.repo.Get()
}

// Create saves new category and returns it
func (srv *CategoryService) Create(name string) (model.Category, error) {
	newCategory := model.Category{
		Name: name,
	}

	if err := srv.repo.Push(&newCategory); err != nil {
		return model.Category{}, err
	}
	return newCategory, nil
}

// Find returns category by ID
func (srv *CategoryService) Find(id uint) (model.Category, error) {
	category := srv.repo.Find(id)

	if category.ID == 0 {
		return model.Category{}, fmt.Errorf(fmt.Sprintf("Category %d not found", id))
	}
	return category, nil
}
