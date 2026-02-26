package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		utils.SendError(w, "Username parameter is required", http.StatusBadRequest)
		return
	}

	if !utils.IsValidUsername(username) {
		utils.SendError(w, "Invalid username format", http.StatusBadRequest)
		return
	}

	var user models.User
	exists := db.DB.Where("username = ?", username).First(&user).Error == nil

	utils.SendSuccess(w, map[string]interface{}{
		"available": !exists,
		"username":  username,
	})
}
