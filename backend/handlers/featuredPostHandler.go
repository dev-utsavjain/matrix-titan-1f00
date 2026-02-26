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

	// Get featured posts, limit to 6 for homepage
	if err := db.DB.Where("featured = ? AND status = ?", true, "published").
		Order("created_at DESC").
		Limit(6).
		Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch featured posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, posts)
}
