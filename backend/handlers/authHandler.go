package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !utils.IsValidUsername(req.Username) {
		utils.SendError(w, "Username must be 3-30 characters and contain no spaces", http.StatusBadRequest)
		return
	}

	if !utils.IsValidEmail(req.Email) {
		utils.SendError(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	if !utils.IsValidPassword(req.Password) {
		utils.SendError(w, "Password must be at least 6 characters", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendError(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		ID:        uuid.NewString(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
	})
}
