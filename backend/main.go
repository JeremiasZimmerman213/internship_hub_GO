package main

import (
	"github.com/gin-gonic/gin"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/config"
)

func main () {
	config.ConnectDB()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})

	r.Run(":8080")
}