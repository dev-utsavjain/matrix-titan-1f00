package handlers

import (
	"encoding/json"
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

// SignupRequest represents the signup request body
type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Signup creates a new user
func Signup(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if !utils.IsValidUsername(req.Username) {
		utils.SendError(w, "Invalid username format", http.StatusBadRequest)
		return
	}
	if !utils.IsValidEmail(req.Email) {
		utils.SendError(w, "Invalid email format", http.StatusBadRequest)
		return
	}
	if len(req.Password) < 6 {
		utils.SendError(w, "Password must be at least 6 characters", http.StatusBadRequest)
		return
	}

	// Check if username exists
	var existingUser models.User
	if err := db.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		utils.SendError(w, "Username already exists", http.StatusConflict)
		return
	}

	// Check if email exists
	if err := db.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		utils.SendError(w, "Email already exists", http.StatusConflict)
		return
	}

	// Create user (password should be hashed in production)
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password, // TODO: Hash password before saving
	}

	if err := db.DB.Create(&user).Error; err != nil {
		utils.SendError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Return user without password
	user.Password = ""
	utils.SendSuccess(w, user)
}
