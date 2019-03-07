package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"nu/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockCategoryService struct{}

func (srv *mockCategoryService) Get() []model.Category {
	categoryList := make([]model.Category, 0)

	for i := 0; i < 3; i++ {
		timeData, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
		categoryList = append(categoryList, model.Category{
			ID:        uint((i + 1)),
			Name:      fmt.Sprintf("Category %d", (i + 1)),
			CreatedAt: timeData,
			UpdatedAt: timeData,
		})
	}

	return categoryList
}

func (srv *mockCategoryService) Create(name string) (model.Category, error) {
	return model.Category{
		ID:        4,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (srv *mockCategoryService) Find(id uint) (model.Category, error) {
	return model.Category{
		ID:        id,
		Name:      _testFaker.Lorem().Word(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (srv *mockCategoryService) Update(id uint, name string) error {
	return nil
}

func (srv *mockCategoryService) Delete(id uint) error {
	return nil
}

func TestCategoryHandlerGet(t *testing.T) {
	categoryHandler := CategoryHandler{
		srv: &mockCategoryService{},
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("CategoryHandler.Get() failed with error")
		t.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(categoryHandler.Get).ServeHTTP(rr, req)

	if status := rr.Code; status != 200 {
		t.Errorf("CategoryHandler.Get() failed with status code %d", status)
	}

	expectedResponse := string(`
		[
			{
				"id": 1,
				"name": "Category 1",
				"createdAt": "2006-01-02 15:04:05",
				"updatedAt": "2006-01-02 15:04:05"
			},
			{
				"id": 2,
				"name": "Category 2",
				"createdAt": "2006-01-02 15:04:05",
				"updatedAt": "2006-01-02 15:04:05"
			},
			{
				"id": 3,
				"name": "Category 3",
				"createdAt": "2006-01-02 15:04:05",
				"updatedAt": "2006-01-02 15:04:05"
			}
		]
	`)

	assert.JSONEq(t, expectedResponse, rr.Body.String())
}
