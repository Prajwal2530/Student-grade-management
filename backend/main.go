package main

import (
	"grade-management-system/config"
	"grade-management-system/models"
	"grade-management-system/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Attempt to load .env file if it exists
	_ = godotenv.Load()

	// 1. Connect to Database
	config.ConnectDatabase()

	// 2. AutoMigrate Models
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Course{},
		&models.Enrollment{},
		&models.Grade{},
	)
	if err != nil {
		log.Fatal("Failed to auto-migrate database schema:", err)
	}
	log.Println("Database schema migrated successfully")

	// 3. Setup Routes
	r := routes.SetupRoutes()

	// 4. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
