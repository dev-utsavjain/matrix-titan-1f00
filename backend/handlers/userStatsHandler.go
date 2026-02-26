package handlers

import (
	"net/http"
	"backend/db"
	"backend/models"
	"backend/utils"
	"backend/views"
)

func GetUserStats(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var totalPosts int64
	if err := db.DB.Model(&models.Post{}).Where("author_id = ?", userID).Count(&totalPosts).Error; err != nil {
		utils.SendError(w, "Failed to count posts", http.StatusInternalServerError)
		return
	}

	var publishedPosts int64
	if err := db.DB.Model(&models.Post{}).Where("author_id = ? AND status = ?", userID, "published").Count(&publishedPosts).Error; err != nil {
		utils.SendError(w, "Failed to count published posts", http.StatusInternalServerError)
		return
	}

	var draftPosts int64
	if err := db.DB.Model(&models.Post{}).Where("author_id = ? AND status = ?", userID, "draft").Count(&draftPosts).Error; err != nil {
		utils.SendError(w, "Failed to count draft posts", http.StatusInternalServerError)
		return
	}

	var totalViews int64
	if err := db.DB.Model(&models.Post{}).Where("author_id = ?", userID).Select("COALESCE(SUM(views), 0)").Scan(&totalViews).Error; err != nil {
		utils.SendError(w, "Failed to sum views", http.StatusInternalServerError)
		return
	}

	var totalComments int64
	if err := db.DB.Model(&models.Comment{}).Joins("JOIN posts ON posts.id = comments.post_id").Where("posts.author_id = ?", userID).Count(&totalComments).Error; err != nil {
		utils.SendError(w, "Failed to count comments", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, views.UserStatsResponse{
		TotalPosts:      int(totalPosts),
		PublishedPosts:  int(publishedPosts),
		DraftPosts:      int(draftPosts),
		TotalViews:      int(totalViews),
		TotalComments:   int(totalComments),
	})
}
