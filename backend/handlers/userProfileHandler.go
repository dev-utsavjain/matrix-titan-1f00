package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

// GetUserProfile returns the current user profile
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	// TODO: Get user ID from authentication context
	// For now, we'll use a placeholder user ID
	userID := "placeholder-user-id"

	var user models.User
	if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		utils.SendError(w, "User not found", http.StatusNotFound)
		return
	}

	// Return user without password
	user.Password = ""
	utils.SendSuccess(w, user)
}
