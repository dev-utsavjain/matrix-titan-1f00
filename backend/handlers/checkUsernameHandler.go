package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"net/http"
)

func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		utils.SendError(w, "Username parameter is required", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		utils.SendSuccess(w, map[string]bool{"available": true})
		return
	}

	utils.SendSuccess(w, map[string]bool{"available": false})
}
