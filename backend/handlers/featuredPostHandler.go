package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

// GetFeaturedPosts returns featured posts for homepage
func GetFeaturedPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post

	if err := db.DB.
		Model(&models.Post{}).
		Where("status = ? AND featured = ?", "published", true).
		Preload("Author").
		Preload("Category").
		Order("created_at DESC").
		Limit(6).
		Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch featured posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, posts)
}