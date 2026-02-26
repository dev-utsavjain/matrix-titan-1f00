package config

import (
	"os"
)

// GetPort returns the port from environment or default
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

// GetAllowedOrigins returns allowed CORS origins
func GetAllowedOrigins() string {
	origins := os.Getenv("ALLOWED_ORIGINS")
	if origins == "" {
		return "*"
	}
	return origins
}
