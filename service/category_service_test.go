package service

import (
	"fmt"
	"nu/model"
	"testing"
)

func (repo *mockCategoryRepository) Construct() {
	repo.categoryList = make([]model.Category, 0)

	for i := 0; i < 5; i++ {
		repo.categoryList = append(repo.categoryList, model.Category{
			ID:   uint((i + 1)),
			Name: _testFaker.Lorem().Word(),
		})
	}
}

func (repo *mockCategoryRepository) Get() []model.Category {
	return repo.categoryList
}

func (repo *mockCategoryRepository) Push(data *model.Category) error {
	newID := func() uint {
		temp := uint(len(repo.categoryList))

		for {
			for _, category := range repo.categoryList {
				if category.ID == temp {
					temp++
					continue
				}
			}
			return temp
		}
	}

	data.ID = newID()
	repo.categoryList = append(repo.categoryList, *data)

	return nil
}

func (repo *mockCategoryRepository) FindBy(field string, value interface{}) model.Category {
	for _, category := range repo.categoryList {
		if field == "id" && category.ID == value {
			return category
		}
		if field == "name" && category.Name == value {
			return category
		}
	}
	return model.Category{}
}

func (repo *mockCategoryRepository) Find(id uint) model.Category {
	return repo.FindBy("id", id)
}

func (repo *mockCategoryRepository) Update(data *model.Category) error {
	for i, category := range repo.categoryList {
		if category.ID == data.ID {
			updatedCat := &repo.categoryList[i]
			updatedCat.Name = data.Name
			return nil
		}
	}
	return fmt.Errorf(fmt.Sprintf("Category %d not found", data.ID))
}

func (repo *mockCategoryRepository) Delete(id uint) error {
	for i, category := range repo.categoryList {
		if category.ID == id {
			repo.categoryList = append(repo.categoryList[:i], repo.categoryList[(i+1):]...)
			return nil
		}
	}
	return fmt.Errorf(fmt.Sprintf("Category %d not found", id))
}

func TestCategoryServiceGet(t *testing.T) {
	categoryRepo := &mockCategoryRepository{}
	categoryRepo.Construct()

	categoryService := &CategoryService{
		repo: categoryRepo,
	}

	categoryList := categoryService.Get()
	if len(categoryList) != 5 {
		t.Errorf("CategoryService.Get() does not return correct len")
	}
}
