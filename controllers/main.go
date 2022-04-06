package controllers

import (
	"github.com/JustSomeHack/go-api-sample/services"

	"gorm.io/gorm"
)

var catsService services.CatsService
var dogsService services.DogsService

func SetupServer(db *gorm.DB) {

}
