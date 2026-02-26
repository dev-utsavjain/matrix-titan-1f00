package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

// GetRelatedPosts returns related posts for a blog post
func GetRelatedPosts(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("id")
	if postID == "" {
		utils.SendError(w, "Post ID is required", http.StatusBadRequest)
		return
	}

	// Get the current post to find its category
	var currentPost models.Post
	if err := db.DB.Where("id = ?", postID).First(&currentPost).Error; err != nil {
		utils.SendError(w, "Post not found", http.StatusNotFound)
		return
	}

	// Get related posts from same category, excluding current post
	var relatedPosts []models.Post
	if err := db.DB.
		Model(&models.Post{}).
		Where("category_id = ? AND id != ? AND status = ?", currentPost.CategoryID, postID, "published").
		Preload("Author").
		Preload("Category").
		Order("created_at DESC").
		Limit(4).
		Find(&relatedPosts).Error; err != nil {
		utils.SendError(w, "Failed to fetch related posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, relatedPosts)
}
