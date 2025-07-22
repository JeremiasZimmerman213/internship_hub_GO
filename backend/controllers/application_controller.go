package controllers

import (
	"net/http"

	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/config"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/models"
	"github.com/gin-gonic/gin"
)

func GetApplications(c *gin.Context) {
	var applications []models.Application
	config.DB.Find(&applications)
	c.JSON(http.StatusOK, applications)
}

func GetApplicationByID(c *gin.Context) {
	id := c.Param("id")

	var application models.Application

	if err := config.DB.First(&application, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	c.JSON(http.StatusOK, application)
}

func CreateApplication(c *gin.Context) {
	var app models.Application
	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Auth is disabled: just set UserID to 1 for now
	app.UserID = 1
	// The following code is commented out for now:
	// // Get the authenticated user from middleware
	// user, exists := c.Get("user")
	// if !exists {
	//     c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
	//     return
	// }
	// // Type assertion to get the user struct
	// authenticatedUser, ok := user.(models.User)
	// if !ok {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data"})
	//     return
	// }
	// // Associate the application with the authenticated user
	// app.UserID = authenticatedUser.ID

	// Create the application with error checking
	if err := config.DB.Create(&app).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, app)
}

func UpdateApplication(c *gin.Context) {
	id := c.Param("id")
	var app models.Application

	if err := config.DB.First(&app, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	// Auth is disabled: skip user context and ownership checks
	// The following code is commented out for now:
	// // Get the authenticated user from middleware
	// user, exists := c.Get("user")
	// if !exists {
	//     c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
	//     return
	// }
	// authenticatedUser, ok := user.(models.User)
	// if !ok {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data"})
	//     return
	// }
	// // Check if the application belongs to the authenticated user
	// if app.UserID != authenticatedUser.ID {
	//     c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own applications"})
	//     return
	// }

	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Auth is disabled: don't overwrite UserID
	// app.UserID = authenticatedUser.ID

	if err := config.DB.Save(&app).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update application: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

func DeleteApplication(c *gin.Context) {
	id := c.Param("id")
	var app models.Application

	if err := config.DB.First(&app, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	// Auth is disabled: skip user context and ownership checks
	// The following code is commented out for now:
	// // Get the authenticated user from middleware
	// user, exists := c.Get("user")
	// if !exists {
	//     c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
	//     return
	// }
	// authenticatedUser, ok := user.(models.User)
	// if !ok {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data"})
	//     return
	// }
	// // Check if the application belongs to the authenticated user
	// if app.UserID != authenticatedUser.ID {
	//     c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own applications"})
	//     return
	// }

	if err := config.DB.Delete(&app).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete application: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}