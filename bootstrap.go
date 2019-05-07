package main

import (
	"github.com/dewadg/nu/db"
	"github.com/dewadg/nu/repository"
	"github.com/dewadg/nu/service"
)

var categoryRepository *repository.CategoryRepository

var categoryService *service.CategoryService

func initRepositories() {
	categoryRepository = repository.NewCategoryRepository(db.Get())
}

func initServices() {
	categoryService = service.NewCategoryService(categoryRepository)
}
