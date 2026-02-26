package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

// GetUserStats returns user dashboard statistics
func GetUserStats(w http.ResponseWriter, r *http.Request) {
	// TODO: Get user ID from authentication context
	// For now, we'll use a placeholder user ID
	userID := "placeholder-user-id"

	var stats struct {
		TotalPosts    int64 `json:"totalPosts"`
		PublishedPosts int64 `json:"publishedPosts"`
		DraftPosts    int64 `json:"draftPosts"`
		TotalViews    int64 `json:"totalViews"`
	}

	// Count total posts
	db.DB.Model(&models.Post{}).Where("author_id = ?", userID).Count(&stats.TotalPosts)

	// Count published posts
	db.DB.Model(&models.Post{}).Where("author_id = ? AND status = ?", userID, "published").Count(&stats.PublishedPosts)

	// Count draft posts
	db.DB.Model(&models.Post{}).Where("author_id = ? AND status = ?", userID, "draft").Count(&stats.DraftPosts)

	// Sum total views
	db.DB.Model(&models.Post{}).Where("author_id = ?", userID).Select("COALESCE(SUM(views), 0)").Row().Scan(&stats.TotalViews)

	utils.SendSuccess(w, stats)
}
