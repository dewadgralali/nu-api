package repository

import (
	"nu/db"
	"nu/model"
	"testing"
)

func TestCategoryRepositoryTest(t *testing.T) {
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
