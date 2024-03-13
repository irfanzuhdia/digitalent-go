package config

import (
	"assignment-2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "test1234"
	dbPort   = "5432"
	dbName   = "assignment-2"
	DB       *gorm.DB
	err      error
)

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, dbPort)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate the schema
	migrateSchema()
}

func migrateSchema() {
	DB.AutoMigrate(&models.Order{}, &models.Item{})
}
