package models

import "time"

type Dog struct {
	Breed     string    `json:"breed"`
	Color     string    `json:"color"`
	Birthdate time.Time `json:"birthdate"`
	Weight    int       `json:"weight"`
}
