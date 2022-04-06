package services

import (
	"os"
	"testing"
	"time"

	"github.com/JustSomeHack/go-api-sample/models"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

var cats []models.Cat
var dogs []models.Dog

func setupTests(t testing.TB) func(t testing.TB) {
	os.Remove("../test.db")

	var err error

	db, err = gorm.Open(sqlite.Open("../test.db"))
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.Cat{}, &models.Dog{}); err != nil {
		panic(err)
	}

	seedData()

	return func(t testing.TB) {
		os.Remove("../test.db")
	}
}

func seedData() {
	cats = []models.Cat{
		{
			ID:        uuid.New(),
			Name:      "Nacho",
			Breed:     "Tabby",
			Color:     "Orange",
			Birthdate: time.Date(2020, 2, 10, 0, 0, 0, 0, time.UTC),
			Weight:    17,
		},
		{
			ID:        uuid.New(),
			Name:      "Captain Marble",
			Breed:     "Main Coon",
			Color:     "Calico",
			Birthdate: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
			Weight:    10,
		},
		{
			ID:        uuid.New(),
			Name:      "Appa",
			Breed:     "Domestic Shorthair",
			Color:     "White/Gray",
			Birthdate: time.Date(2020, 9, 19, 0, 0, 0, 0, time.UTC),
			Weight:    19,
		},
	}

	err := db.Create(&cats).Error
	if err != nil {
		panic(err)
	}

	dogs = []models.Dog{
		{
			ID:        uuid.New(),
			Name:      "Snowball",
			Breed:     "Shiba Inu",
			Color:     "Cream",
			Birthdate: time.Date(2020, 6, 12, 0, 0, 0, 0, time.UTC),
			Weight:    22,
		},
		{
			ID:        uuid.New(),
			Name:      "Koda",
			Breed:     "Samoyed",
			Color:     "White",
			Birthdate: time.Date(2022, 1, 28, 0, 0, 0, 0, time.UTC),
			Weight:    5,
		},
	}

	err = db.Create(&dogs).Error
	if err != nil {
		panic(err)
	}
}
