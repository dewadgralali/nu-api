package service

import (
	"testing"

	"github.com/jaswdr/faker"
)

var _testFaker faker.Faker

func TestMain(m *testing.M) {
	_testFaker = faker.New()

	m.Run()
}
