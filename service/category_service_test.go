package service

import (
	"reflect"
	"testing"

	localMock "github.com/dewadg/nu-api/mock"
	"github.com/dewadg/nu-api/model"
)

func TestCategoryServiceGet(t *testing.T) {
	categoryRepo := &localMock.CategoryRepositoryMock{}
	categoryService := &CategoryService{
		repo: categoryRepo,
	}

	expected := []model.Category{
		model.Category{
			ID:   1,
			Name: _testFaker.Person().Name(),
		},
		model.Category{
			ID:   2,
			Name: _testFaker.Person().Name(),
		},
		model.Category{
			ID:   3,
			Name: _testFaker.Person().Name(),
		},
	}
	categoryRepo.On("Get").Return(expected)

	actual := categoryService.Get()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("CategoryService.Get() does not return expected data")
	}
}

func TestCategoryServiceCreate(t *testing.T) {
	categoryRepo := &localMock.CategoryRepositoryMock{}
	categoryService := &CategoryService{
		repo: categoryRepo,
	}

	name := _testFaker.Lorem().Word()
	expected := model.Category{
		ID:   0,
		Name: name,
	}
	payload := model.Category{
		Name: name,
	}
	categoryRepo.On("Push", &payload).Return(nil)

	actual, err := categoryService.Create(name)
	if err != nil {
		t.Errorf("CategoryService.Create() failed with error")
		t.Errorf(err.Error())
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("CategoryService.Create() failed to save data")
	}
}

func TestCategoryServiceFind(t *testing.T) {
	categoryRepo := &localMock.CategoryRepositoryMock{}
	categoryService := &CategoryService{
		repo: categoryRepo,
	}

	expected := model.Category{
		ID:   1,
		Name: _testFaker.Person().Name(),
	}
	payload := uint(1)
	categoryRepo.On("Find", payload).Return(expected)

	actual, err := categoryService.Find(payload)
	if err != nil {
		t.Errorf("CategoryService.Find() failed with error")
		t.Errorf(err.Error())
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("CategoryService.Find() failed to return the correct data")
	}
}

func TestCategoryServiceUpdate(t *testing.T) {
	categoryRepo := &localMock.CategoryRepositoryMock{}
	categoryService := &CategoryService{
		repo: categoryRepo,
	}

	payload := model.Category{
		ID:   uint(1),
		Name: _testFaker.Person().Name(),
	}
	categoryRepo.On("Update", &payload).Return(nil)

	err := categoryService.Update(payload.ID, payload.Name)
	if err != nil {
		t.Errorf("CategoryService.Update() failed with error")
		t.Errorf(err.Error())
	}
}

func TestCategoryServiceDelete(t *testing.T) {
	categoryRepo := &localMock.CategoryRepositoryMock{}
	categoryService := &CategoryService{
		repo: categoryRepo,
	}

	payload := uint(1)
	categoryRepo.On("Delete", payload).Return(nil)

	err := categoryService.Delete(payload)
	if err != nil {
		t.Errorf("CategoryService.Delete() failed with error")
		t.Errorf(err.Error())
	}
}
