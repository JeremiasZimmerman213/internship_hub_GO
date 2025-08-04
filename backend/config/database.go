package config

import (
	"fmt"
	"log"
	"os"

	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func ConnectDB() {
	host := getEnvOrDefault("DB_HOST", "localhost")
	port := getEnvOrDefault("DB_PORT", "5433")
	user := getEnvOrDefault("DB_USER", "tracker_user")
	password := getEnvOrDefault("DB_PASSWORD", "tracker_pass")
	db_name := getEnvOrDefault("DB_NAME", "internship_tracker")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, db_name,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Connected to PostgreSQL DB successfully.")
	DB = database

	DB.AutoMigrate(&models.Application{}, &models.User{})
}
