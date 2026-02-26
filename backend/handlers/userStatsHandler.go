package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

type UserStatsResponse struct {
	TotalPosts    int64 `json:"totalPosts"`
	PublishedPosts int64 `json:"publishedPosts"`
	DraftPosts    int64 `json:"draftPosts"`
	TotalViews    int64 `json:"totalViews"`
}

func GetUserStats(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var stats UserStatsResponse

	db.DB.Model(&models.Post{}).Where("author_id = ?", userID).Count(&stats.TotalPosts)
	db.DB.Model(&models.Post{}).Where("author_id = ? AND status = ?", userID, "published").Count(&stats.PublishedPosts)
	db.DB.Model(&models.Post{}).Where("author_id = ? AND status = ?", userID, "draft").Count(&stats.DraftPosts)

	var views sql.NullInt64
	db.DB.Model(&models.Post{}).Where("author_id = ?", userID).Select("COALESCE(SUM(views), 0)").Row().Scan(&views)
	if views.Valid {
		stats.TotalViews = views.Int64
	}

	utils.SendSuccess(w, stats)
}
