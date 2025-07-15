package utils

import (
	"time"
	// "os"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("super-secret-key") // In production, use os.Getenv()

func GenerateJWT(username string) (string, error) {
	// Create the claims (payload)
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 3 days
	}

	// Create the token using HS256 algorithm and sign it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}