package repository

import (
	"fmt"
	"github.com/dewadg/nu-api/db"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var _testFaker faker.Faker
var _testDB *gorm.DB

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env.test")
	if err != nil {
		fmt.Println("No .env.test file was specified")
	}

	_testDB = db.Get()
	_testFaker = faker.New()

	m.Run()
}
