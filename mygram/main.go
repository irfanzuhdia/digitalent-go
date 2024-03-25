package main

import (
	"mygram/app"
	"mygram/config"
	"os"
)

func main() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "test1234")
	os.Setenv("DB_NAME", "mygram")
	os.Setenv("DB_PORT", "5432")

	os.Setenv("JWT_SECRET", "your_jwt_secret")

	db := config.SetupDatabaseConnection()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	app.StartServer(db)
}
