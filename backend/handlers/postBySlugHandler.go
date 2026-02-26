package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

// GetPostBySlug returns a single post by slug
func GetPostBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		utils.SendError(w, "Slug is required", http.StatusBadRequest)
		return
	}

	var post models.Post
	if err := db.DB.
		Model(&models.Post{}).
		Where("slug = ? AND status = ?", slug, "published").
		Preload("Author").
		Preload("Category").
		Preload("Comments.Author").
		First(&post).Error; err != nil {
		utils.SendError(w, "Post not found", http.StatusNotFound)
		return
	}

	// Increment view count
	db.DB.Model(&post).Update("views", post.Views+1)

	utils.SendSuccess(w, post)
}
