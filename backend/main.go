package main

import (
	"log"
	"os"

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

	// Add CORS middleware (allow all origins for now)
	r.Use(cors.Default())

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
        protected.DELETE("/applications/:id", controllers.DeleteApplication)
	}

	r.Static("/uploads", "./uploads")

	port := getEnvOrDefault("PORT", "8080")
	r.Run(":" + port)
}
