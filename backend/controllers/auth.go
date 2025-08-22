package controllers

import (
	"net/http"
	"os"
	"time"
	"strings"
	
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/services"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/config"
	"github.com/JeremiasZimmerman213/internship_hub_GO/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var authInput models.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	var userFound models.User
    config.DB.Where("username=?", authInput.Username).Find(&userFound)
    if userFound.ID != 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
        return
    }

	config.DB.Where("email=?", authInput.Email).Find(&userFound)
	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	user := models.User{
        Username:   authInput.Username,
        Email:      authInput.Email,
        Password:   string(passwordHash),
        IsVerified: false,
    }

	if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

	// Generate verification token
    token := services.GenerateVerificationToken()
    verification := models.EmailVerification{
        UserID:    user.ID,
        Token:     token,
        ExpiresAt: time.Now().Add(24 * time.Hour),
    }

	if err := config.DB.Create(&verification).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create verification token"})
        return
    }

	if err := services.SendVerificationEmail(user.Email, token); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
        return
    }

	c.JSON(http.StatusOK, gin.H{
        "message": "User created successfully. Please check your email for verification.",
        "user_id": user.ID,
    })
}

func Login(c *gin.Context) {
	var loginInput models.LoginInput

	if err := c.ShouldBindJSON(&loginInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	var userFound models.User
	if strings.Contains(loginInput.UsernameOrEmail, "@") {
        // It's an email
        config.DB.Where("email=?", loginInput.UsernameOrEmail).Find(&userFound)
    } else {
        // It's a username
        config.DB.Where("username=?", loginInput.UsernameOrEmail).Find(&userFound)
    }

	if userFound.ID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
        return
    }

	if !userFound.IsVerified {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Please verify your email address before logging in",
            "needs_verification": true,
        })
        return
    }

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(loginInput.Password)); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
        return
    }

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":  userFound.ID,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    })

	token, err := generateToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
        return
    }

	c.JSON(200, gin.H{
        "token": token,
        "user": gin.H{
            "id": userFound.ID,
            "username": userFound.Username,
            "email": userFound.Email,
        },
    })
}

func VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Verification token is required"})
		return
	}

	var verification models.EmailVerification
    if err := config.DB.Where("token = ? AND expires_at > ?", token, time.Now()).First(&verification).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired verification token"})
        return
    }

	var user models.User
    if err := config.DB.First(&user, verification.UserID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

	user.IsVerified = true
    if err := config.DB.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user"})
        return
    }

	config.DB.Delete(&verification)

	c.JSON(http.StatusOK, gin.H{
        "message": "Email verified successfully! You can now log in.",
    })
}

func ResendVerification(c *gin.Context) {
	var input struct {
        Email string `json:"email" binding:"required,email"`
    }

	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if user.IsVerified {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email already verified"})
        return
    }

    // Delete old verification tokens
    config.DB.Where("user_id = ?", user.ID).Delete(&models.EmailVerification{})

    // Generate new token
    token := services.GenerateVerificationToken()
    verification := models.EmailVerification{
        UserID:    user.ID,
        Token:     token,
        ExpiresAt: time.Now().Add(24 * time.Hour),
    }
    
    config.DB.Create(&verification)

    // Send new verification email
    if err := services.SendVerificationEmail(user.Email, token); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Verification email sent successfully",
    })
}

func GetUserProfile(c *gin.Context) {
	user, _ := c.Get("currentUser")
	c.JSON(200, gin.H{
		"user": user,
	})
}
