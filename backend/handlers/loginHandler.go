package handlers

import (
	"encoding/json"
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

// LoginRequest represents the login request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login authenticates a user
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Username == "" || req.Password == "" {
		utils.SendError(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Find user by username
	var user models.User
	if err := db.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		utils.SendError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Verify password (TODO: Implement proper password hashing)
	if user.Password != req.Password {
		utils.SendError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Return user without password
	user.Password = ""
	utils.SendSuccess(w, user)
}
