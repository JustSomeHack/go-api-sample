package controllers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/one-byte-data/go-api-sample/internal/middlewares"
	"github.com/one-byte-data/go-api-sample/internal/services"

	"gorm.io/gorm"
)

var catsService services.CatsService
var dogsService services.DogsService

func SetupRouter(db *gorm.DB) (*gin.Engine, error) {
	catsService = services.NewCatsService(db)
	dogsService = services.NewDogsService(db)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	router.Use(middlewares.ValidateHeader())

	health := router.Group("/health")
	{
		health.GET("", HealthGet)
	}

	cats := router.Group("/cats")
	{
		cats.DELETE("/:id", CatsDelete)
		cats.POST("/count", CatsCount)
		cats.GET("", CatsGet)
		cats.GET("/:id", CatsGetOne)
		cats.POST("", CatsPost)
		cats.PUT("/:id", CatsPut)
	}

	dogs := router.Group("/dogs")
	{
		dogs.DELETE("/:id", DogsDelete)
		dogs.POST("/count", DogsCount)
		dogs.GET("", DogsGet)
		dogs.GET("/:id", DogsGetOne)
		dogs.POST("", DogsPost)
		dogs.PUT("/:id", DogsPut)
	}

	return router, nil
}
