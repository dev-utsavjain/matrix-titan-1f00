package handlers

import (
	"net/http"
	"strings"

	"backend/db"
	"backend/models"
	"backend/utils"
	"gorm.io/gorm"
)

// GetUserProfile returns the current user profile
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by auth middleware)
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		utils.SendError(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	var user models.User
	if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.SendError(w, "User not found", http.StatusNotFound)
			return
		}
		utils.SendError(w, "Failed to fetch user profile", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, user)
}

// GetUserPosts returns posts created by the current user
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by auth middleware)
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		utils.SendError(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	var posts []models.Post
	if err := db.DB.Where("author_id = ?", userID).Order("created_at DESC").Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch user posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, posts)
}

// GetUserStats returns user dashboard statistics
func GetUserStats(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by auth middleware)
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		utils.SendError(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	stats := make(map[string]interface{})

	// Total posts count
	var totalPosts int64
	if err := db.DB.Model(&models.Post{}).Where("author_id = ?", userID).Count(&totalPosts).Error; err != nil {
		utils.SendError(w, "Failed to fetch post count", http.StatusInternalServerError)
		return
	}
	stats["totalPosts"] = totalPosts

	// Published posts count
	var publishedPosts int64
	if err := db.DB.Model(&models.Post{}).Where("author_id = ? AND status = ?", userID, "published").Count(&publishedPosts).Error; err != nil {
		utils.SendError(w, "Failed to fetch published post count", http.StatusInternalServerError)
		return
	}
	stats["publishedPosts"] = publishedPosts

	// Draft posts count
	var draftPosts int64
	if err := db.DB.Model(&models.Post{}).Where("author_id = ? AND status = ?", userID, "draft").Count(&draftPosts).Error; err != nil {
		utils.SendError(w, "Failed to fetch draft post count", http.StatusInternalServerError)
		return
	}
	stats["draftPosts"] = draftPosts

	// Total views count
	var totalViews int64
	if err := db.DB.Model(&models.Post{}).Where("author_id = ?", userID).Select("COALESCE(SUM(views), 0)").Scan(&totalViews).Error; err != nil {
		stats["totalViews"] = 0
	} else {
		stats["totalViews"] = totalViews
	}

	// Recent posts (last 30 days)
	var recentPosts int64
	if err := db.DB.Model(&models.Post{}).
		Where("author_id = ? AND created_at >= ?", userID, "now() - interval '30 days'").
		Count(&recentPosts).Error; err != nil {
		stats["recentPosts"] = 0
	} else {
		stats["recentPosts"] = recentPosts
	}

	utils.SendSuccess(w, stats)
}
