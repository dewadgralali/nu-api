package service

import (
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
