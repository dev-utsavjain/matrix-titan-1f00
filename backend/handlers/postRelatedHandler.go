package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

func GetRelatedPosts(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.SendError(w, "Post ID is required", http.StatusBadRequest)
		return
	}

	var post models.Post
	if err := db.DB.Where("id = ? AND status = ?", id, "published").First(&post).Error; err != nil {
		utils.SendError(w, "Post not found", http.StatusNotFound)
		return
	}

	var related []models.Post
	if err := db.DB.Where("id != ? AND status = ? AND category_id = ?", post.ID, "published", post.CategoryID).
		Preload("Author").
		Preload("Category").
		Order("created_at DESC").
		Limit(4).
		Find(&related).Error; err != nil {
		utils.SendError(w, "Failed to fetch related posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, related)
}
