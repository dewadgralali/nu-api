package mock

import (
	"github.com/dewadg/nu-api/model"
	"github.com/stretchr/testify/mock"
)

// PostRepositoryMock mocks PostRepository
type PostRepositoryMock struct {
	mock.Mock
}

// Get mocks Get
func (r *PostRepositoryMock) Get() []model.Post {
	args := r.Called()

	return args.Get(0).([]model.Post)
}

// Push mocks Push
func (r *PostRepositoryMock) Push(data *model.Post) error {
	args := r.Called(data)

	return args.Error(0)
}

// FindBy mocks FindBy
func (r *PostRepositoryMock) FindBy(name string, value interface{}) model.Post {
	args := r.Called(name, value)

	return args.Get(0).(model.Post)
}

// Find mocks Find
func (r *PostRepositoryMock) Find(id uint) model.Post {
	args := r.Called(id)

	return args.Get(0).(model.Post)
}

// Update mocks Update
func (r *PostRepositoryMock) Update(data *model.Post) error {
	args := r.Called(data)

	return args.Error(0)
}

// Delete mocks Delete
func (r *PostRepositoryMock) Delete(id uint) error {
	args := r.Called(id)

	return args.Error(0)
}
