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

func getCurrentUser(c *gin.Context) (*models.User, error) {
    userInterface, exists := c.Get("currentUser")
    if !exists {
        return nil, fmt.Errorf("user not found in context")
    }
    
    user, ok := userInterface.(models.User)
    if !ok {
        return nil, fmt.Errorf("invalid user type in context")
    }
    
    return &user, nil
}

func GetApplications(c *gin.Context) {
    user, err := getCurrentUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var applications []models.Application
    config.DB.Where("user_id = ?", user.ID).Order("applied_date DESC, company ASC").Find(&applications)
    c.JSON(http.StatusOK, applications)
}

func GetApplicationByID(c *gin.Context) {
    user, err := getCurrentUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    id := c.Param("id")
    var application models.Application

    if err := config.DB.Where("id = ? AND user_id = ?", id, user.ID).First(&application).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
        return
    }

    c.JSON(http.StatusOK, application)
}

func CreateApplication(c *gin.Context) {
    user, err := getCurrentUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

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
    uploadPath := "./uploads/" + filename

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

    // Parse status as integer (ApplicationStatus)
    statusInt := c.PostForm("status")
    var status models.ApplicationStatus
    if statusInt == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
        return
    }
    // Convert string to uint8
    var statusVal uint8
    _, err = fmt.Sscanf(statusInt, "%d", &statusVal)
    if err != nil || statusVal > uint8(models.StatusRejected) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})
        return
    }
    status = models.ApplicationStatus(statusVal)

    app := models.Application{
        Company:     c.PostForm("company"),
        Position:    c.PostForm("position"),
        Status:      status,
        Location:    c.PostForm("location"),
        AppliedDate: appliedDate,
        Term:        c.PostForm("term"),
        Note:        c.PostForm("note"),
        ResumeURL:   "/uploads/" + filename,
        UserID:      user.ID, // Use authenticated user's ID
    }

    if err := config.DB.Create(&app).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application: " + err.Error()})
        return
    }

    c.JSON(http.StatusCreated, app)
}

func DeleteApplication(c *gin.Context) {
    user, err := getCurrentUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    id := c.Param("id")
    var app models.Application

    if err := config.DB.Where("id = ? AND user_id = ?", id, user.ID).First(&app).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
        return
    }

    // Try to delete the file first before deleting from database
    if app.ResumeURL != "" {
        filePath := "." + app.ResumeURL

        fmt.Printf("DEBUG: Attempting to delete file at path: %s\n", filePath)
        fmt.Printf("DEBUG: Resume URL from DB: %s\n", app.ResumeURL)

        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            fmt.Printf("DEBUG: File does not exist at path: %s\n", filePath)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Resume file not found on server: " + filePath})
            return
        }

        if err := os.Remove(filePath); err != nil {
            fmt.Printf("DEBUG: Failed to delete file %s: %v\n", filePath, err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete resume file: " + err.Error()})
            return
        }

        fmt.Printf("DEBUG: Successfully deleted file: %s\n", filePath)
    }

    if err := config.DB.Delete(&app).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete application: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}

func UpdateApplication(c *gin.Context) {
    user, err := getCurrentUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    id := c.Param("id")
    var app models.Application

    // Find the existing application (user-specific)
    if err := config.DB.Where("id = ? AND user_id = ?", id, user.ID).First(&app).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
        return
    }

    c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 5<<20)

    if err := c.Request.ParseMultipartForm(5 << 20); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Request body too large (Max 5MB)"})
        return
    }

    // Handle optional file upload (new resume)
    file, fileHeader, err := c.Request.FormFile("resume")
    var newResumeURL string
    if err == nil {
        // New file uploaded, validate and save it
        defer file.Close()

        if fileHeader.Header.Get("Content-Type") != "application/pdf" &&
            fileHeader.Filename[len(fileHeader.Filename)-4:] != ".pdf" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Only PDF files are allowed"})
            return
        }

        filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), fileHeader.Filename)
        uploadPath := "./uploads/" + filename

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

        // Delete old resume file if it exists
        if app.ResumeURL != "" {
            oldFilePath := "." + app.ResumeURL
            if err := os.Remove(oldFilePath); err != nil {
                fmt.Printf("Warning: Failed to delete old resume file %s: %v\n", oldFilePath, err)
            }
        }

        newResumeURL = "/uploads/" + filename
    } else {
        // No new file uploaded, keep existing resume URL
        newResumeURL = app.ResumeURL
    }

    // Parse applied date
    appliedDate, err := time.Parse(time.RFC3339, c.PostForm("applied_date"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
        return
    }

    // Parse status
    statusInt := c.PostForm("status")
    var status models.ApplicationStatus
    if statusInt == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
        return
    }
    var statusVal uint8
    _, err = fmt.Sscanf(statusInt, "%d", &statusVal)
    if err != nil || statusVal > uint8(models.StatusRejected) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})
        return
    }
    status = models.ApplicationStatus(statusVal)

    // Update the application
    app.Company = c.PostForm("company")
    app.Position = c.PostForm("position")
    app.Status = status
    app.Location = c.PostForm("location")
    app.AppliedDate = appliedDate
    app.Term = c.PostForm("term")
    app.Note = c.PostForm("note")
    app.ResumeURL = newResumeURL

    if err := config.DB.Save(&app).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update application: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, app)
}
