package controllers

import (
	"io"
	"net/http"

	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/config"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/models"
	"github.com/gin-gonic/gin"

	"fmt"
	"os"
	"time"
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
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 5<<20)

	if err := c.Request.ParseMultipartForm(5 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body too large (Max 5MB)"})
		return
	}

	file, fileHeader, err := c.Request.FormFile("resume")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Resume PDF is required"})
		return
	}

	defer file.Close()

	if fileHeader.Header.Get("Content-Type") != "application/pdf" &&
	fileHeader.Filename[len(fileHeader.Filename)-4:] != ".pdf" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only PDF files are allowed"})
		return
	}

	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), fileHeader.Filename)
	uploadPath := "../uploads/" + filename

	out, err := os.Create(uploadPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file: " + err.Error()})
		return
	}

	defer out.Close()
	if _, err := io.Copy(out, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file: " + err.Error()})
		return
	}
	
	appliedDate, err := time.Parse(time.RFC3339, c.PostForm("applied_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	app := models.Application{
		Company:    c.PostForm("company"),
		Position:   c.PostForm("position"),
		Status:    c.PostForm("status"),
		Location:  c.PostForm("location"),
		AppliedDate: appliedDate,
		ResumeURL: "/uploads/" + filename,
		UserID: 1, // Hardcoded for now, replace with authenticated user ID later
	}

	if err := config.DB.Create(&app).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, app)
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