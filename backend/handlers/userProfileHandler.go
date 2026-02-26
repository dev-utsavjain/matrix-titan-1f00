package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"net/http"
)

func GetCurrentUserProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var user models.User
	if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		utils.SendError(w, "User not found", http.StatusNotFound)
		return
	}

	utils.SendSuccess(w, user)
}
