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
