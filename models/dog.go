package models

import (
	"time"

	"github.com/google/uuid"
)

type Dog struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" gorm:"check:name <> ''"`
	Breed     string    `json:"breed" gorm:"check:breed <> ''"`
	Color     string    `json:"color" gorm:"check:color <> ''"`
	Birthdate time.Time `json:"birthdate"`
	Weight    int       `json:"weight" gorm:"check:weight > 0"`
}
