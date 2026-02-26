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

	// Get recent posts, limit to 10 for homepage
	if err := db.DB.Where("status = ?", "published").
		Order("created_at DESC").
		Limit(10).
		Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch recent posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, posts)
}
