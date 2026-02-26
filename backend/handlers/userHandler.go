package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"backend/views"
	"context"
	"net/http"
	"github.com/google/uuid"
)

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value("user_id").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		utils.SendError(w, "Invalid user ID", http.StatusBadRequest)
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
	userIDStr := r.Context().Value("user_id").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		utils.SendError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var posts []models.Post
	if err := db.DB.Where("author_id = ?", userID).Find(&posts).Error; err != nil {
		utils.SendError(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, posts)
}

func GetUserStats(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value("user_id").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		utils.SendError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var totalPosts int64
	var publishedPosts int64
	var draftPosts int64
	var totalViews int64

	db.DB.Model(&models.Post{}).Where("author_id = ?", userID).Count(&totalPosts)
	db.DB.Model(&models.Post{}).Where("author_id = ? AND status = ?", userID, "published").Count(&publishedPosts)
	db.DB.Model(&models.Post{}).Where("author_id = ? AND status = ?", userID, "draft").Count(&draftPosts)
	db.DB.Model(&models.Post{}).Where("author_id = ?", userID).Select("COALESCE(SUM(views), 0)").Scan(&totalViews)

	stats := views.UserStatsResponse{
		TotalPosts:     totalPosts,
		PublishedPosts: publishedPosts,
		DraftPosts:     draftPosts,
		TotalViews:     totalViews,
	}

	utils.SendSuccess(w, stats)
}
