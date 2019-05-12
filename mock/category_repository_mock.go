package mock

import (
	"github.com/dewadg/nu/model"
	"github.com/stretchr/testify/mock"
)

// CategoryRepositoryMock mocks CategoryRepository
type CategoryRepositoryMock struct {
	mock.Mock
}

// Get mocks Get
func (r *CategoryRepositoryMock) Get() []model.Category {
	args := r.Called()

	return args.Get(0).([]model.Category)
}

// Push mocks Push
func (r *CategoryRepositoryMock) Push(data *model.Category) error {
	args := r.Called(data)

	return args.Error(0)
}

// FindBy mocks FindBy
func (r *CategoryRepositoryMock) FindBy(field string, value interface{}) model.Category {
	args := r.Called(field, value)

	return args.Get(0).(model.Category)
}

// Find mocks Find
func (r *CategoryRepositoryMock) Find(id uint) model.Category {
	args := r.Called(id)

	return args.Get(0).(model.Category)
}

// Update mocks Update
func (r *CategoryRepositoryMock) Update(data *model.Category) error {
	args := r.Called(data)

	return args.Error(0)
}

// Delete mocks Delete
func (r *CategoryRepositoryMock) Delete(id uint) error {
	args := r.Called(id)

	return args.Error(0)
}
