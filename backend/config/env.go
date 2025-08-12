package config

import (
	"fmt"
	"log"
	"os"
)

func GetJWTSecret() string {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        log.Println("Warning: JWT_SECRET not set, using default (not secure for production)")
        return "default-jwt-secret-change-in-production"
    }
    return secret
}

// ValidateEnv checks that all required environment variables are set
func ValidateEnv() {
	requiredEnvVars := []string{
		"JWT_SECRET",
	}

	var missingVars []string
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			missingVars = append(missingVars, envVar)
		}
	}

	if len(missingVars) > 0 {
		log.Printf("Warning: The following environment variables are not set: %v", missingVars)
		log.Println("Using default values for development. Set these in production!")
	}
}

// PrintConfig logs the current configuration (without sensitive data)
func PrintConfig() {
	fmt.Println("=== Application Configuration ===")
	fmt.Printf("DB_HOST: %s\n", getEnvOrDefault("DB_HOST", "localhost"))
	fmt.Printf("DB_PORT: %s\n", getEnvOrDefault("DB_PORT", "5433"))
	fmt.Printf("DB_NAME: %s\n", getEnvOrDefault("DB_NAME", "internship_tracker"))
	fmt.Printf("PORT: %s\n", getEnvOrDefault("PORT", "8080"))
	fmt.Printf("GIN_MODE: %s\n", getEnvOrDefault("GIN_MODE", "debug"))
	fmt.Printf("JWT_SECRET: %s\n", func() string {
		if os.Getenv("JWT_SECRET") != "" {
			return "***SET***"
		}
		return "***USING_DEFAULT***"
	}())
	fmt.Println("=====================================")
}
