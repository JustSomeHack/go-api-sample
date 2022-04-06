package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/JustSomeHack/go-api-sample/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

var Cats []models.Cat
var Dogs []models.Dog

func SetupTests(t testing.TB, dialector gorm.Dialector) func(t testing.TB) {
	if _, err := os.Stat("../test.db"); err == nil {
		os.Remove("../test.db")
	}

	var err error

	DB, err = gorm.Open(dialector)
	if err != nil {
		panic(err)
	}

	if err := DB.AutoMigrate(&models.Cat{}, &models.Dog{}); err != nil {
		panic(err)
	}

	LoadCats()
	if err := DB.Create(Cats).Error; err != nil {
		panic(err)
	}

	LoadDogs()
	if err := DB.Create(Dogs).Error; err != nil {
		panic(err)
	}

	return func(t testing.TB) {
		DB.Migrator().DropTable(&models.Cat{})
		DB.Migrator().DropTable(&models.Dog{})

		if _, err := os.Stat("../test.db"); err == nil {
			os.Remove("../test.db")
		}
	}
}

func LoadCats() {
	Cats = make([]models.Cat, 0)
	data, err := os.ReadFile("../tests/cats.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &Cats)
	if err != nil {
		panic(err)
	}
}

func LoadDogs() {
	Dogs = make([]models.Dog, 0)
	data, err := os.ReadFile("../tests/dogs.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &Dogs)
	if err != nil {
		panic(err)
	}
}
