package repository

import (
	"fmt"
	"nu/db"
	"nu/model"
	"testing"
)

func generateCategories() {
	for i := 0; i < 3; i++ {
		_testDB.Create(&model.Category{
			Name: _testFaker.Lorem().Word(),
		})
	}
}

func generatePosts() {
	categoryRepository := &CategoryRepository{
		db: _testDB,
	}

	for i := 0; i < 5; i++ {
		_testDB.Create(&model.Post{
			Title:      fmt.Sprintf("Hello %d", (i + 1)),
			Slug:       "hello",
			MDContent:  "## Hello",
			Categories: categoryRepository.Get(),
		})
	}
}

func TestPostRepositoryGet(t *testing.T) {
	db.Reset()
	generateCategories()
	generatePosts()

	postRepository := &PostRepository{
		db: _testDB,
	}

	postList := postRepository.Get()
	if len(postList) != 5 {
		t.Errorf("PostRepository.Get() doesn't return expected len 5")
	}
}
