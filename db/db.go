package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// Postgres Dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// DB is for handling connection with database
	DB *gorm.DB
)

// Connect to database
func Connect() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, username, dbName, password) //Build connection string
	fmt.Println(dbURI)
	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Success")
	DB = conn
}
