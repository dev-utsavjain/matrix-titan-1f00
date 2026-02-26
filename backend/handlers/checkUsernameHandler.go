package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

// CheckUsername checks if a username is available
func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		utils.SendError(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Validate username format
	if !utils.IsValidUsername(username) {
		utils.SendError(w, "Invalid username format", http.StatusBadRequest)
		return
	}

	// Check if username exists
	var user models.User
	exists := db.DB.Where("username = ?", username).First(&user).Error == nil

	response := map[string]interface{}{
		"available": !exists,
		"username":  username,
	}

	utils.SendSuccess(w, response)
}
