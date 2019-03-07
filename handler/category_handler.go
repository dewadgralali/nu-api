package handler

import (
	"fmt"
	"net/http"
	"nu/object"
	"nu/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// CategoryHandler is sets of category-related
// HTTP handlers
type CategoryHandler struct {
	srv service.CategoryServiceContract
}

// GetRoutes returns this handler routes.
func (hndlr *CategoryHandler) GetRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", hndlr.Get)
	r.Post("/", hndlr.Store)

	return r
}

// Get doc
func (hndlr *CategoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	categoryListResponse := object.CreateCategoryListResponse(hndlr.srv.Get())

	if err := render.RenderList(w, r, categoryListResponse); err != nil {
		fmt.Println(err.Error())
		return
	}
}

// Store doc
func (hndlr *CategoryHandler) Store(w http.ResponseWriter, r *http.Request) {
	payload := object.StoreCategoryRequest{}
	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, createUnprocessableEntityResponse(err.Error()))
		return
	}

	category, err := hndlr.srv.Create(payload.Name)
	if err != nil {
		render.Render(w, r, createInternalServerErrorResponse(err.Error()))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, object.CreateCategoryResponse(category))
}
