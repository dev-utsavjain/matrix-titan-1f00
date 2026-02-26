package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	if err := db.DB.Where("status = ?", "published").
		Order("created_at DESC").
		Preload("Author").
		Preload("Category").
		Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, posts)
}