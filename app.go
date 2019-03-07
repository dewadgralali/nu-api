package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nu/db"
	"nu/handler"
	"nu/repository"
	"nu/service"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

const version = 1

var categoryRepository *repository.CategoryRepository

var categoryService *service.CategoryService

func initRepositories() {
	categoryRepository = repository.NewCategoryRepository(db.Get())
}

func initServices() {
	categoryService = service.NewCategoryService(categoryRepository)
}

func createRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := map[string]interface{}{
			"name":    "Nu API",
			"version": version,
		}
		res, _ := json.Marshal(payload)
		w.Write(res)
	}))

	r.Mount("/categories", (handler.NewCategoryHandler(categoryService)).GetRoutes())

	return r
}

func serveHTTP() {
	router := createRouter()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("App running on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

func runMigration() {
	fmt.Println("Running migrations...")
	db.Migrate()
	fmt.Println("Migrated successfully")
}
