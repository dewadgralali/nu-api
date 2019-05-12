package service

import "github.com/dewadg/nu-api/repository"

// NewCategoryService returns new CategoryService instance
func NewCategoryService(repo repository.CategoryRepositoryContract) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}
