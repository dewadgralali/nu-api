package service

import (
	"nu/model"
	"testing"

	"github.com/jaswdr/faker"
)

var _testFaker faker.Faker

type mockCategoryRepository struct {
	categoryList []model.Category
}

func TestMain(m *testing.M) {
	_testFaker = faker.New()

	m.Run()
}
