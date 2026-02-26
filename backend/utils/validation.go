package utils

import (
	"strings"
)

func IsValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func IsValidUsername(username string) bool {
	return len(username) >= 3 && len(username) <= 30 && strings.TrimSpace(username) == username
}

func IsValidPassword(password string) bool {
	return len(password) >= 6
}
