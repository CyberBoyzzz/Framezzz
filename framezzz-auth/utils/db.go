package utils

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB initializes and returns the database connection
func InitDB() *gorm.DB {
	var err error
	for i := 0; i < 5; i++ {
		// Retry logic: Try to connect to the database
		db, err = gorm.Open(postgres.Open(getDSN()), &gorm.Config{})
		if err == nil {
			log.Println("Database connected successfully")
			return db
		}

		// Wait for 5 seconds before retrying
		log.Printf("Error connecting to the database: %v. Retrying in 5 seconds...\n", err)
		time.Sleep(5 * time.Second)
	}

	log.Fatalf("Failed to connect to the database after 5 retries: %v", err)
	return nil
}

// getDSN returns the Data Source Name (DSN) for the database connection
func getDSN() string {
	return "host=postgres user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
}
