package handler

import "github.com/dewadg/nu-api/service"

// NewCategoryHandler returns new instance
// of CategoryHandler
func NewCategoryHandler(srv *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		srv: srv,
	}
}
