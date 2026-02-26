package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"net/http"
)

func CheckUsernameAvailability(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		utils.SendError(w, "Username is required", http.StatusBadRequest)
		return
	}

	username = utils.SanitizeString(username)
	if !utils.IsValidUsername(username) {
		utils.SendError(w, "Invalid username format", http.StatusBadRequest)
		return
	}

	var count int64
	db.DB.Model(&models.User{}).Where("username = ?", username).Count(&count)

	utils.SendSuccess(w, map[string]interface{}{
		"available": count == 0,
	})
}
