package handlers

import (
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

type ProfileResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar,omitempty"`
	Bio       string `json:"bio,omitempty"`
	CreatedAt string `json:"createdAt"`
}

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

	response := ProfileResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}

	utils.SendSuccess(w, response)
}
