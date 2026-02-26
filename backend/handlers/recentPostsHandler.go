package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

// GetRecentPosts returns recent posts for homepage
func GetRecentPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post

	if err := db.DB.
		Model(&models.Post{}).
		Where("status = ?", "published").
		Preload("Author").
		Preload("Category").
		Order("created_at DESC").
		Limit(6).
		Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch recent posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, posts)
}
