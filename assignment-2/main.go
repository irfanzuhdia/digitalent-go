package main

import (
	"assignment-2/config"
	"assignment-2/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	config.ConnectDB()

	// Initialize Gin router
	router := gin.Default()

	// Define routes
	handlers.SetupRoutes(router)

	// Start server
	router.Run(":8080")
}
