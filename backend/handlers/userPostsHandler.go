package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var posts []models.Post
	if err := db.DB.Where("author_id = ?", userID).
		Preload("Category").
		Order("created_at DESC").
		Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch user posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, posts)
}
