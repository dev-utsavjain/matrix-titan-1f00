package handlers

import (
	"net/http"
	"strconv"

	"backend/db"
	"backend/models"
	"backend/utils"
	"github.com/google/uuid"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit

	categorySlug := r.URL.Query().Get("category")
	status := r.URL.Query().Get("status")

	query := db.DB.Model(&models.Post{}).Preload("Author").Preload("Category")

	if status == "" {
		status = "published"
	}
	query = query.Where("status = ?", status)

	if categorySlug != "" {
		var category models.Category
		if err := db.DB.Where("slug = ?", categorySlug).First(&category).Error; err == nil {
			query = query.Where("category_id = ?", category.ID)
		}
	}

	var total int64
	query.Count(&total)

	var posts []models.Post
	if err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, map[string]interface{}{
		"posts": posts,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func GetFeaturedPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	if err := db.DB.Preload("Author").Preload("Category").
		Where("featured = ? AND status = ?", true, "published").
		Order("created_at DESC").
		Limit(6).
		Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch featured posts", http.StatusInternalServerError)
		return
	}
	utils.SendSuccess(w, posts)
}

func GetRecentPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	if err := db.DB.Preload("Author").Preload("Category").
		Where("status = ?", "published").
		Order("created_at DESC").
		Limit(6).
		Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch recent posts", http.StatusInternalServerError)
		return
	}
	utils.SendSuccess(w, posts)
}

func GetPostBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		utils.SendError(w, "Slug is required", http.StatusBadRequest)
		return
	}

	var post models.Post
	if err := db.DB.Preload("Author").Preload("Category").
		Where("slug = ? AND status = ?", slug, "published").
		First(&post).Error; err != nil {
		utils.SendError(w, "Post not found", http.StatusNotFound)
		return
	}

	// Increment views
	db.DB.Model(&post).Update("views", post.Views+1)

	utils.SendSuccess(w, post)
}

func GetRelatedPosts(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		utils.SendError(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.SendError(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var post models.Post
	if err := db.DB.Where("id = ?", id).First(&post).Error; err != nil {
		utils.SendError(w, "Post not found", http.StatusNotFound)
		return
	}

	var related []models.Post
	query := db.DB.Preload("Author").Preload("Category").
		Where("id != ? AND status = ?", id, "published")

	if post.CategoryID != uuid.Nil {
		query = query.Where("category_id = ?", post.CategoryID)
	}

	if err := query.Order("created_at DESC").Limit(4).Find(&related).Error; err != nil {
		utils.SendError(w, "Failed to fetch related posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, related)
}
