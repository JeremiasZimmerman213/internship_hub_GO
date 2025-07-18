package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// For development only - in production, this should be a fatal error
		secret = "development-secret-key-change-in-production"
	}
	return []byte(secret)
}

func GenerateJWT(username string) (string, error) {
	if username == "" {
		return "", errors.New("username cannot be empty")
	}

	// Create the claims (payload)
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 3 days
		"iat":      time.Now().Unix(),                     // issued at
	}

	// Create the token using HS256 algorithm and sign it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(getJWTSecret())
}
