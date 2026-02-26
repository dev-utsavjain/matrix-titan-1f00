package utils

import (
	"regexp"
	"strings"
)

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(strings.TrimSpace(email))
}

func IsValidUsername(username string) bool {
	if len(username) < 3 || len(username) > 30 {
		return false
	}
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return usernameRegex.MatchString(username)
}

func SanitizeString(s string) string {
	return strings.TrimSpace(s)
}
