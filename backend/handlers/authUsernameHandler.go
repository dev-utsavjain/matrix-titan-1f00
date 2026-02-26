package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
	"backend/views"
)

func CheckUsernameAvailability(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		utils.SendError(w, "Username is required", http.StatusBadRequest)
		return
	}

	if !utils.IsValidUsername(username) {
		utils.SendError(w, "Invalid username format", http.StatusBadRequest)
		return
	}

	var existingUser models.User
	available := true
	if err := db.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		available = false
	}

	response := views.UsernameCheckResponse{
		Available: available,
		Username:  username,
	}

	utils.SendSuccess(w, response)
}
