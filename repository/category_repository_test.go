package repository

import (
	"nu/db"
	"nu/model"
	"testing"
)

func TestCategoryRepositoryGet(t *testing.T) {
	db.Reset()

	for i := 0; i < 5; i++ {
		_testDB.Create(&model.Category{
			Name: _testFaker.Lorem().Word(),
		})
	}

	categoryRepo := &CategoryRepository{
		db: _testDB,
	}
	categoryList := categoryRepo.Get()

	if len(categoryList) != 5 {
		t.Errorf("CategoryRepository.Get() doesn't return expected len 5")
	}
}

func TestCategoryRepositoryPush(t *testing.T) {
	db.Reset()
	categoryRepo := &CategoryRepository{
		db: _testDB,
	}

	newCategory := model.Category{
		Name: _testFaker.Lorem().Word(),
	}
	err := categoryRepo.Push(&newCategory)

	if err != nil {
		t.Errorf("CategoryRepository.Push() failed with error")
		t.Errorf(err.Error())
	}
	if newCategory.ID == 0 {
		t.Errorf("CategoryRepository.Push() failed to generate ID")
	}
}