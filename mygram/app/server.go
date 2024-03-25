package app

import (
	"log"

	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) {
	server := InitializeServer(db)

	log.Printf("Server started on port 8080")
	err := server.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
