package service

import "github.com/dewadg/nu-api/model"

// CategoryServiceContract represents interface for
// category service
type CategoryServiceContract interface {
	Get() []model.Category
	Create(name string) (model.Category, error)
	Find(id uint) (model.Category, error)
	Update(id uint, name string) error
	Delete(id uint) error
}

// PostServiceContract represents interface for
// post service
type PostServiceContract interface {
	Get() []model.Post
	Create(title string, content string, categoriesID []uint) (model.Post, error)
	Find(id uint) (model.Post, error)
	FindBySlug(slug string) (model.Post, error)
	Update(id uint, title string, content string, categoriesID []uint) (model.Post, error)
	Delete(id uint) error
}
