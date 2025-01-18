package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/CyberBoyzzz/Framezzz/routes"
	"github.com/CyberBoyzzz/Framezzz/utils" 
	"github.com/CyberBoyzzz/Framezzz/models" 
)

func main() {
	// Initialize DB connection
	db := utils.InitDB()

	// Run database migrations (if applicable)
	// This can be done after initializing the DB connection, such as:
	db.AutoMigrate(&models.User{}) // Auto migrate User model, for example

	// Initialize Gin router
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	// Start the server on port 7312
	if err := router.Run(":7313"); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
