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
	if err := db.DB.Where("id = ? AND status = ?", postID, "published").First(&currentPost).Error; err != nil {
		if err.Error() == "record not found" {
			utils.SendError(w, "Post not found", http.StatusNotFound)
			return
		}
		utils.SendError(w, "Failed to fetch post", http.StatusInternalServerError)
		return
	}

	// Get related posts (same category, different post, published)
	var relatedPosts []models.Post
	if err := db.DB.Where("category_id = ? AND id != ? AND status = ?", 
		currentPost.CategoryID, currentPost.ID, "published").
		Order("created_at DESC").
		Limit(5).
		Find(&relatedPosts).Error; err != nil {
		utils.SendError(w, "Failed to fetch related posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, relatedPosts)
}
