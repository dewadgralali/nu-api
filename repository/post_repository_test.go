package repository

import (
	"fmt"
	"github.com/dewadg/nu/db"
	"github.com/dewadg/nu/model"
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
		t.Errorf("PostRepository.FindBy() failed to find post by ID")
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
		t.Errorf("PostRepository.Find() failed to find post by ID")
	}
}

func TestRepositoryUpdate(t *testing.T) {
	db.Reset()
	generateCategories()
	generatePosts()

	postRepository := &PostRepository{
		db: _testDB,
	}

	updatedPost := postRepository.Find(3)
	updatedPost.Title = "Hellooo"
	updatedPost.Slug = "Hellooo"
	updatedPost.MDContent = "Hellooo"

	err := postRepository.Update(&updatedPost)
	if err != nil {
		t.Errorf("PostRepository.Update() failed with error")
		t.Errorf(err.Error())
	}

	expectedPost := postRepository.Find(3)
	if expectedPost.Title != "Hellooo" || expectedPost.Slug != "Hellooo" || expectedPost.MDContent != "Hellooo" {
		t.Errorf("PostRepository.Update() failed to update post")
	}
}

func TestRepositoryDelete(t *testing.T) {
	db.Reset()
	generateCategories()
	generatePosts()

	postRepository := &PostRepository{
		db: _testDB,
	}

	err := postRepository.Delete(3)
	if err != nil {
		t.Errorf("PostRepository.Delete() failed with error")
		t.Errorf(err.Error())
	}

	expectedPost := postRepository.Find(3)
	if expectedPost.ID != 0 {
		t.Errorf("PostRepository.Delete() failed to remove post")
	}
}
