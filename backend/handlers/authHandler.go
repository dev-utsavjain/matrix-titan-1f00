package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"backend/views"
	"encoding/json"
	"net/http"
	"backend/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var req views.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req.Username = utils.SanitizeString(req.Username)
	req.Email = utils.SanitizeString(req.Email)
	req.Password = utils.SanitizeString(req.Password)

	if !utils.IsValidUsername(req.Username) {
		utils.SendError(w, "Invalid username format", http.StatusBadRequest)
		return
	}

	if !utils.IsValidEmail(req.Email) {
		utils.SendError(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	if len(req.Password) < 6 {
		utils.SendError(w, "Password must be at least 6 characters", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.SendError(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		utils.SendError(w, "Failed to create user", http.StatusInternalServerError)
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

func Login(w http.ResponseWriter, r *http.Request) {
	var req views.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req.Email = utils.SanitizeString(req.Email)
	req.Password = utils.SanitizeString(req.Password)

	var user models.User
	if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		utils.SendError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		utils.SendError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.ID.String())
	if err != nil {
		utils.SendError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, map[string]interface{}{
		"token": token,
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"avatar":   user.Avatar,
			"bio":      user.Bio,
		},
	})
}
