package models

import "time"

type Cat struct {
	Breed     string    `json:"breed"`
	Color     string    `json:"color"`
	Birthdate time.Time `json:"birthdate"`
	Weight    int       `json:"weight"`
}
