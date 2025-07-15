package main

import (
	"github.com/gin-gonic/gin"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/config"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/controllers"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/handlers"
)

func main () {
	config.ConnectDB()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})

	r.GET("/applications", controllers.GetApplications)
	r.GET("/applications/:id", controllers.GetApplicationByID)
	r.POST("/applications", controllers.CreateApplication)
	r.PUT("/applications/:id", controllers.UpdateApplication)
	r.DELETE("/applications/:id", controllers.DeleteApplication)

	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", controllers.Login)

	r.Run(":8080")
}