package models

import (
	"time"

	"github.com/google/uuid"
)

type Dog struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" binding:"required,min=2,max=24" gorm:"check:name <> ''"`
	Breed     string    `json:"breed" binding:"required,min=2,max=24" gorm:"check:breed <> ''"`
	Color     string    `json:"color" binding:"required,min=2,max=24" gorm:"check:color <> ''"`
	Birthdate time.Time `json:"birthdate" binding:"required"`
	Weight    int       `json:"weight" binding:"required,gte=1,lt=300" gorm:"check:weight > 0"`
}
