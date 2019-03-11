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

func TestPostRepositoryPush(t *testing.T) {
	db.Reset()
	generateCategories()

	categoryRepository := &CategoryRepository{
		db: _testDB,
	}
	postRepository := &PostRepository{
		db: _testDB,
	}

	newPost := model.Post{
		Title:      _testFaker.Lorem().Word(),
		Slug:       _testFaker.Lorem().Word(),
		MDContent:  _testFaker.Lorem().Word(),
		Categories: categoryRepository.Get(),
	}

	err := postRepository.Push(&newPost)

	if err != nil {
		t.Errorf("PostRepository.Push() failed with error")
		t.Errorf(err.Error())
	}
	if newPost.ID == 0 {
		t.Errorf("PostRepository.Push() failed to generate ID")
	}
}

func TestPostRepositoryFindBy(t *testing.T) {
	db.Reset()
	generateCategories()
	generatePosts()

	postRepository := &PostRepository{
		db: _testDB,
	}

	post := postRepository.FindBy("id", 1)

	if post.ID != 1 {
		t.Errorf("CategoryRepository.FindBy() failed to find post by ID")
	}
}

func TestPostRepositoryFind(t *testing.T) {
	db.Reset()
	generateCategories()
	generatePosts()

	postRepository := &PostRepository{
		db: _testDB,
	}

	post := postRepository.Find(1)

	if post.ID != 1 {
		t.Errorf("CategoryRepository.Find() failed to find post by ID")
	}
}
