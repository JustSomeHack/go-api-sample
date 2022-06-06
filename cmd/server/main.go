package main

import (
	"fmt"
	"os"

	_ "github.com/one-byte-data/go-api-sample/docs"
	"github.com/one-byte-data/go-api-sample/internal/controllers"
	"github.com/one-byte-data/go-api-sample/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var version string = "development"

// @title Go API Sample
// @version 1.0
// @description This is a sample API in go
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	printVersion()

	db := setupDatabase(getConnectionString())

	router, err := controllers.SetupRouter(db)
	if err != nil {
		panic(err)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(); err != nil {
		panic(err)
	}
}

func getConnectionString() string {
	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		panic("unable to get a connection string")
	}
	return connectionString
}

func printVersion() {
	fmt.Printf("Starting go-api-sample %s\n\n", version)
}

func setupDatabase(connectionString string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic(fmt.Sprintf("unable to connect to database: %v", err))
	}

	if err := db.AutoMigrate(&models.Cat{}, &models.Dog{}); err != nil {
		panic(err)
	}

	return db
}
