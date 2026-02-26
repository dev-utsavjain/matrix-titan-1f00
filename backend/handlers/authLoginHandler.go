package handlers

import (
	"encoding/json"
	"net/http"

	"backend/db"
	"backend/models"
	"backend/utils"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar,omitempty"`
	Bio      string `json:"bio,omitempty"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.DB.Where("username = ? OR email = ?", req.Username, req.Username).First(&user).Error; err != nil {
		utils.SendError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		utils.SendError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	response := LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Bio:      user.Bio,
	}

	utils.SendSuccess(w, response)
}
