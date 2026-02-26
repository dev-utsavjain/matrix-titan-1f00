package handlers

import (
	"net/http"
	"strconv"

	"backend/db"
	"backend/models"
	"backend/utils"
)

// GetPosts returns all published posts with pagination and filtering
func GetPosts(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	category := r.URL.Query().Get("category")
	author := r.URL.Query().Get("author")

	// Default pagination values
	page := 1
	limit := 10

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	offset := (page - 1) * limit

	// Build query
	query := db.DB.Model(&models.Post{}).Where("status = ?", "published")

	if category != "" {
		query = query.Where("category_id = ?", category)
	}
	if author != "" {
		query = query.Where("author_id = ?", author)
	}

	// Get total count
	var total int64
	if err := query.Count(&total).Error; err != nil {
		utils.SendError(w, "Failed to count posts", http.StatusInternalServerError)
		return
	}

	// Get posts with pagination
	var posts []models.Post
	if err := query.
		Preload("Author").
		Preload("Category").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := map[string]interface{}{
		"posts": posts,
		"pagination": map[string]interface{}{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": (total + int64(limit) - 1) / int64(limit),
		},
	}

	utils.SendSuccess(w, response)
}
