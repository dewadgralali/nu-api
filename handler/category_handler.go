package handler

import (
	"context"
	"fmt"
	"net/http"
	"nu/model"
	"nu/object"
	"nu/service"
	"strconv"

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
	r.Route("/{categoryID}", func(r chi.Router) {
		r.Use(hndlr.Context)
		r.Get("/", hndlr.GetOne)
		r.Patch("/", hndlr.Update)
	})

	return r
}

// Context of these routes.
func (hndlr *CategoryHandler) Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categoryID, _ := strconv.Atoi(chi.URLParam(r, "categoryID"))
		category, err := hndlr.srv.Find(uint(categoryID))
		if err != nil {
			render.Render(w, r, createNotFoundResponse(err.Error()))
			return
		}

		ctx := context.WithValue(r.Context(), categoryCtx, category)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
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

// GetOne doc
func (hndlr *CategoryHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category, ok := ctx.Value(categoryCtx).(model.Category)
	if !ok {
		render.Render(w, r, createUnprocessableEntityResponse(""))
		return
	}

	render.Render(w, r, object.CreateCategoryResponse(category))
}

// Update doc
func (hndlr *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category, ok := ctx.Value(categoryCtx).(model.Category)
	if !ok {
		render.Render(w, r, createUnprocessableEntityResponse(""))
		return
	}

	payload := object.UpdateCategoryRequest{}
	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, createUnprocessableEntityResponse(err.Error()))
		return
	}

	err := hndlr.srv.Update(category.ID, payload.Name)
	if err != nil {
		render.Render(w, r, createInternalServerErrorResponse(err.Error()))
		return
	}

	updatedCategory, _ := hndlr.srv.Find(category.ID)

	render.Render(w, r, object.CreateCategoryResponse(updatedCategory))
}
