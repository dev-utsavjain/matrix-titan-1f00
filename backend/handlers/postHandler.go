package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"net/http"
	"strconv"
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

	category := r.URL.Query().Get("category")
	author := r.URL.Query().Get("author")

	query := db.DB.Model(&models.Post{}).Where("status = ?", "published")

	if category != "" {
		query = query.Where("category_id = ?", category)
	}
	if author != "" {
		query = query.Where("author_id = ?", author)
	}

	var posts []models.Post
	var total int64

	if err := query.Count(&total).Error; err != nil {
		utils.SendError(w, "Failed to count posts", http.StatusInternalServerError)
		return
	}

	if err := query.Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, map[string]interface{}{
		"posts": posts,
		"pagination": map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

func GetFeaturedPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	if err := db.DB.Where("status = ? AND featured = ?", "published", true).Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch featured posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, posts)
}

func GetRecentPosts(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 || limit > 20 {
		limit = 5
	}

	var posts []models.Post
	if err := db.DB.Where("status = ?", "published").Order("created_at DESC").Limit(limit).Find(&posts).Error; err != nil {
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
	if err := db.DB.Where("slug = ? AND status = ?", slug, "published").First(&post).Error; err != nil {
		utils.SendError(w, "Post not found", http.StatusNotFound)
		return
	}

	db.DB.Model(&post).Update("views", post.Views+1)

	utils.SendSuccess(w, post)
}

func GetRelatedPosts(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.SendError(w, "ID is required", http.StatusBadRequest)
		return
	}

	var post models.Post
	if err := db.DB.Where("id = ? AND status = ?", id, "published").First(&post).Error; err != nil {
		utils.SendError(w, "Post not found", http.StatusNotFound)
		return
	}

	var related []models.Post
	if err := db.DB.Where("id != ? AND category_id = ? AND status = ?", post.ID, post.CategoryID, "published").Limit(3).Find(&related).Error; err != nil {
		utils.SendError(w, "Failed to fetch related posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, related)
}
