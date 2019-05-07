package repository

import "github.com/dewadg/nu/model"

// CategoryRepositoryContract represents interface for
// category repository
type CategoryRepositoryContract interface {
	Get() []model.Category
	Push(data *model.Category) error
	FindBy(field string, value interface{}) model.Category
	Find(id uint) model.Category
	Update(data *model.Category) error
	Delete(id uint) error
}

// PostRepositoryContract represents interface for
// post repository
type PostRepositoryContract interface {
	Get() []model.Post
	Push(data *model.Post) error
	FindBy(field string, value interface{}) model.Post
	Find(id uint) model.Post
	Update(data *model.Post) error
	Delete(id uint) error
}
