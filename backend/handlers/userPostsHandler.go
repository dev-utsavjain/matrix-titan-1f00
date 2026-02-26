package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

// GetUserPosts returns posts created by the current user
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	// TODO: Get user ID from authentication context
	// For now, we'll use a placeholder user ID
	userID := "placeholder-user-id"

	var posts []models.Post
	if err := db.DB.
		Model(&models.Post{}).
		Where("author_id = ?", userID).
		Preload("Category").
		Order("created_at DESC").
		Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch user posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, posts)
}
