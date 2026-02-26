package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"net/http"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	if err := db.DB.Find(&categories).Error; err != nil {
		utils.SendError(w, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}
	utils.SendSuccess(w, categories)
}
