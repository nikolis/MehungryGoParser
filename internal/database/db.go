package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Global variable to hold the DB connection
var DB *gorm.DB

// Initialize the database connection
func InitDB(dsn string) {
	var err error
	// Open the connection to the PostgreSQL database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connection established")
}

// Automigrate database schema (optional, can be used to create tables)
func AutoMigrate(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}
}
