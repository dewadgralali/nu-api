package object

import (
	"errors"
	"net/http"
)

// StoreCategoryRequest represents request object
// for storing new category
type StoreCategoryRequest struct {
	Name string `json:"name"`
}

// Bind validates the request
func (req *StoreCategoryRequest) Bind(r *http.Request) error {
	if req.Name == "" {
		return errors.New("`name` cannot be empty")
	}
	return nil
}
