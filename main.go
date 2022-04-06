package main

import (
	"fmt"
	"os"

	"github.com/JustSomeHack/go-api-sample/controllers"
	"github.com/JustSomeHack/go-api-sample/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var version string = "development"

func main() {
	printVersion()

	db := setupDatabase(getConnectionString())

	router, err := controllers.SetupRouter(db)
	if err != nil {
		panic(err)
	}

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
