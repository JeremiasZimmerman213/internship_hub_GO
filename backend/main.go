package main

import (
	"log"
	"os"

	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/config"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/controllers"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/handlers"

	// "github.com/JeremiasZimmerman213/internship_hub_GO/backend/middleware"
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

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/applications", controllers.GetApplications)
	r.GET("/applications/:id", controllers.GetApplicationByID)
	// r.POST("/applications", controllers.CreateApplication)
	// r.PUT("/applications/:id", controllers.UpdateApplication)
	// r.DELETE("/applications/:id", controllers.DeleteApplication)

	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", controllers.Login)

	r.POST("/applications", controllers.CreateApplication)
	r.PUT("/applications/:id", controllers.UpdateApplication)
	r.DELETE("/applications/:id", controllers.DeleteApplication)

	// Auth middleware and protected routes are disabled for now
	// auth := r.Group("/")
	// auth.Use(middleware.AuthMiddleware())
	// {
	//     auth.POST("/applications", controllers.CreateApplication)
	//     auth.PUT("/applications/:id", controllers.UpdateApplication)
	//     auth.DELETE("/applications/:id", controllers.DeleteApplication)
	// }

	r.Static("/uploads", "./uploads")

	port := getEnvOrDefault("PORT", "8080")
	r.Run(":" + port)
}
