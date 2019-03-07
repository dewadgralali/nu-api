package main

import (
	"nu/db"
	"nu/repository"
	"nu/service"
)

var categoryRepository *repository.CategoryRepository

var categoryService *service.CategoryService

func initRepositories() {
	categoryRepository = repository.NewCategoryRepository(db.Get())
}

func initServices() {
	categoryService = service.NewCategoryService(categoryRepository)
}
