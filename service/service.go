package service

import "nu/repository"

// NewCategoryService returns new CategoryService instance
func NewCategoryService(repo repository.CategoryRepositoryContract) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}
