package main

import (
	"fmt"
	"os"

	"github.com/JustSomeHack/go-api-sample/controllers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var version string = "development"

func main() {
	printVersion()

	db := setupDatabase(getConnectionString())

	controllers.SetupServer(db)
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
	return db
}
