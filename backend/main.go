package main

import (
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/config"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/controllers"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/handlers"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

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

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/applications", controllers.CreateApplication)
		auth.PUT("/applications/:id", controllers.UpdateApplication)
		auth.DELETE("/applications/:id", controllers.DeleteApplication)
	}

	r.Run(":8080")
}
