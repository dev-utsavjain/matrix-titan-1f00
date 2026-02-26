package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateToken generates a random token
func GenerateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
