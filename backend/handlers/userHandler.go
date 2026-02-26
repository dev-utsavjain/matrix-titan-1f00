package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var user models.User
	if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		utils.SendError(w, "User not found", http.StatusNotFound)
		return
	}

	utils.SendSuccess(w, map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"avatar":   user.Avatar,
		"bio":      user.Bio,
		"createdAt": user.CreatedAt,
	})
}

func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var posts []models.Post
	if err := db.DB.Preload("Category").Where("author_id = ?", userID).Order("created_at DESC").Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch user posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, posts)
}

func GetUserStats(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var totalPosts int64
	db.DB.Model(&models.Post{}).Where("author_id = ?", userID).Count(&totalPosts)

	var publishedPosts int64
	db.DB.Model(&models.Post{}).Where("author_id = ? AND status = ?", userID, "published").Count(&publishedPosts)

	var draftPosts int64
	db.DB.Model(&models.Post{}).Where("author_id = ? AND status = ?", userID, "draft").Count(&draftPosts)

	var totalViews int64
	db.DB.Model(&models.Post{}).Where("author_id = ?", userID).Select("COALESCE(SUM(views), 0)").Scan(&totalViews)

	utils.SendSuccess(w, map[string]interface{}{
		"totalPosts":     totalPosts,
		"publishedPosts": publishedPosts,
		"draftPosts":     draftPosts,
		"totalViews":     totalViews,
	})
}
