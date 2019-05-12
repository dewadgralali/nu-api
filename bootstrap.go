package main

import (
	"github.com/dewadg/nu-api/db"
	"github.com/dewadg/nu-api/repository"
	"github.com/dewadg/nu-api/service"
)

var categoryRepository *repository.CategoryRepository

var categoryService *service.CategoryService

func initRepositories() {
	categoryRepository = repository.NewCategoryRepository(db.Get())
}

func initServices() {
	categoryService = service.NewCategoryService(categoryRepository)
}
