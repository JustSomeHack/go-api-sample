package models

import (
	"time"

	"github.com/google/uuid"
)

type Dog struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Breed     string    `json:"breed"`
	Color     string    `json:"color"`
	Birthdate time.Time `json:"birthdate"`
	Weight    int       `json:"weight"`
}
