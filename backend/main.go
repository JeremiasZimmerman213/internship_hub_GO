package main

import (
	"log"
	"os"
	"strings"

	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/config"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/controllers"

	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	// Load .env file if it exists (for development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Validate environment configuration
	config.ValidateEnv()
	config.PrintConfig()

	config.ConnectDB()

	r := gin.Default()

	// Add CORS middleware with proper configuration
	allowedOrigins := []string{"http://localhost:5173", "http://localhost:3000"}

	// Add additional origins from environment variable for network access
	if extraOrigins := os.Getenv("CORS_ALLOWED_ORIGINS"); extraOrigins != "" {
		// Split by comma and add to allowed origins
		origins := strings.Split(extraOrigins, ",")
		for _, origin := range origins {
			allowedOrigins = append(allowedOrigins, strings.TrimSpace(origin))
		}
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/auth/signup", controllers.CreateUser)
	r.POST("/auth/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(middleware.CheckAuth)
	{
		// User profile
		protected.GET("/user/profile", controllers.GetUserProfile)

		// Application routes - all protected and user-specific
		protected.GET("/applications", controllers.GetApplications)
		protected.GET("/applications/:id", controllers.GetApplicationByID)
		protected.POST("/applications", controllers.CreateApplication)
		protected.PUT("/applications/:id", controllers.UpdateApplication)
		protected.PATCH("/applications/:id/status", controllers.UpdateApplicationStatus)
		protected.DELETE("/applications/:id", controllers.DeleteApplication)
	}

	r.Static("/uploads", "./uploads")

	port := getEnvOrDefault("PORT", "8080")
	// r.Run(":" + port)
	r.Run("0.0.0.0:" + port)
}
