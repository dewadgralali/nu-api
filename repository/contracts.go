package repository

import "nu/model"

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
