package utils

import (
	"regexp"
	"strings"
)

// IsValidEmail checks if email format is valid
func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(strings.TrimSpace(email))
}

// IsValidUsername checks if username is valid (alphanumeric and underscores, 3-20 characters)
func IsValidUsername(username string) bool {
	if len(username) < 3 || len(username) > 20 {
		return false
	}
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return usernameRegex.MatchString(username)
}

// SanitizeString removes leading/trailing whitespace
func SanitizeString(s string) string {
	return strings.TrimSpace(s)
}
