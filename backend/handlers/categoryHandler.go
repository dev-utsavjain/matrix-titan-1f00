package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	if err := db.DB.Order("name ASC").Find(&categories).Error; err != nil {
		utils.SendError(w, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}
	utils.SendSuccess(w, categories)
}
