package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"github.com/dewadg/nu/model"
	"strconv"
	"testing"
	"time"

	"github.com/go-chi/chi"

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
	timeData, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	return model.Category{
		ID:        1,
		Name:      name,
		CreatedAt: timeData,
		UpdatedAt: timeData,
	}, nil
}

func (srv *mockCategoryService) Find(id uint) (model.Category, error) {
	timeData, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	return model.Category{
		ID:        id,
		Name:      "Category 1",
		CreatedAt: timeData,
		UpdatedAt: timeData,
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

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	http.HandlerFunc(categoryHandler.Get).ServeHTTP(w, r)

	if status := w.Code; status != 200 {
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

	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestCategoryHandlerStore(t *testing.T) {
	categoryHandler := CategoryHandler{
		srv: &mockCategoryService{},
	}
	payload := map[string]string{
		"name": "New Category",
	}
	reqData, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", bytes.NewBuffer(reqData))
	r.Header.Set("Content-Type", "application/json")

	http.HandlerFunc(categoryHandler.Store).ServeHTTP(w, r)

	if status := w.Code; status != http.StatusCreated {
		t.Errorf("CategoryService.Store() failed with status code %d", status)
	}

	expectedResponse := string(`
		{
			"id": 1,
			"name": "New Category",
			"createdAt": "2006-01-02 15:04:05",
			"updatedAt": "2006-01-02 15:04:05"
		}
	`)

	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestCategoryHandlerGetOne(t *testing.T) {
	expectedID := 1

	categoryHandler := CategoryHandler{
		srv: &mockCategoryService{},
	}

	handler := http.HandlerFunc(categoryHandler.GetOne)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	rCtx := chi.NewRouteContext()
	rCtx.URLParams.Add("categoryID", strconv.Itoa(expectedID))
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rCtx))

	categoryHandler.Context(handler).ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("CategoryHandler.GetOne() failed with status %d", status)
	}

	expectedResponse := string(`
		{
			"id": 1,
			"name": "Category 1",
			"createdAt": "2006-01-02 15:04:05",
			"updatedAt": "2006-01-02 15:04:05"
		}
	`)

	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestCategoryHandlerUpdate(t *testing.T) {
	expectedID := 1

	categoryHandler := CategoryHandler{
		srv: &mockCategoryService{},
	}

	payload := map[string]string{
		"name": "Category 1",
	}
	reqData, _ := json.Marshal(payload)

	handler := http.HandlerFunc(categoryHandler.Update)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PATCH", "/", bytes.NewBuffer(reqData))

	rCtx := chi.NewRouteContext()
	rCtx.URLParams.Add("categoryID", strconv.Itoa(expectedID))
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rCtx))
	r.Header.Set("Content-Type", "application/json")

	categoryHandler.Context(handler).ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("CategoryHandler.GetOne() failed with status %d", status)
	}

	expectedResponse := string(`
		{
			"id": 1,
			"name": "Category 1",
			"createdAt": "2006-01-02 15:04:05",
			"updatedAt": "2006-01-02 15:04:05"
		}
	`)

	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestCategoryHandlerDestroy(t *testing.T) {
	expectedID := 1

	categoryHandler := CategoryHandler{
		srv: &mockCategoryService{},
	}

	handler := http.HandlerFunc(categoryHandler.Destroy)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)

	rCtx := chi.NewRouteContext()
	rCtx.URLParams.Add("categoryID", strconv.Itoa(expectedID))
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rCtx))

	categoryHandler.Context(handler).ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("CategoryHandler.GetOne() failed with status %d", status)
	}
}
