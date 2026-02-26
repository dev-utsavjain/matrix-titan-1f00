package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"net/http"
)

func GetRelatedPosts(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("id")
	if postID == "" {
		utils.SendError(w, "Post ID is required", http.StatusBadRequest)
		return
	}

	var post models.Post
	if err := db.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		utils.SendError(w, "Post not found", http.StatusNotFound)
		return
	}

	var relatedPosts []models.Post
	query := db.DB.Where("id != ? AND status = ? AND category_id = ?", postID, "published", post.CategoryID).
		Order("created_at DESC").
		Limit(5)

	if err := query.Find(&relatedPosts).Error; err != nil {
		utils.SendError(w, "Failed to fetch related posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, relatedPosts)
}
