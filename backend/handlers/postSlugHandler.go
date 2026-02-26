package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

func GetPostBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		utils.SendError(w, "Slug is required", http.StatusBadRequest)
		return
	}

	var post models.Post
	if err := db.DB.Where("slug = ? AND status = ?", slug, "published").
		Preload("Author").
		Preload("Category").
		First(&post).Error; err != nil {
		utils.SendError(w, "Post not found", http.StatusNotFound)
		return
	}

	db.DB.Model(&post).Update("views", post.Views+1)

	utils.SendSuccess(w, post)
}
