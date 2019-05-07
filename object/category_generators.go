package object

import (
	"github.com/dewadg/nu/model"

	"github.com/go-chi/render"
)

// CreateCategoryResponse returns CategoryResponse
// based on Category model
func CreateCategoryResponse(data model.Category) render.Renderer {
	return &CategoryResponse{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// CreateCategoryListResponse returns slice of CategoryResponse
func CreateCategoryListResponse(data []model.Category) []render.Renderer {
	payload := make([]render.Renderer, 0)

	for _, category := range data {
		payload = append(payload, CreateCategoryResponse(category))
	}
	return payload
}
