package service

import "nu/model"

// CategoryServiceContract represents interface for
// category service
type CategoryServiceContract interface {
	Get() []model.Category
	Create(name string) (model.Category, error)
	Find(id uint) (model.Category, error)
	Update(id uint, name string) error
	Delete(id uint) error
}
