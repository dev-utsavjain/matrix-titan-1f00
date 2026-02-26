package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"backend/db"
	"backend/models"
	"backend/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req.Username = utils.SanitizeString(req.Username)
	req.Email = utils.SanitizeString(req.Email)

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

	var existingUser models.User
	if err := db.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		utils.SendError(w, "Username already exists", http.StatusConflict)
		return
	}

	if err := db.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		utils.SendError(w, "Email already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendError(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		ID:       uuid.New(),
		Username: req.Username,
		Email:    strings.ToLower(req.Email),
		Password: string(hashedPassword),
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

func Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		utils.SendError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.SendError(w, "Invalid credentials", http.StatusUnauthorized)
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

func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		utils.SendError(w, "Username is required", http.StatusBadRequest)
		return
	}

	if !utils.IsValidUsername(username) {
		utils.SendError(w, "Invalid username format", http.StatusBadRequest)
		return
	}

	var count int64
	db.DB.Model(&models.User{}).Where("username = ?", username).Count(&count)

	utils.SendSuccess(w, map[string]interface{}{
		"available": count == 0,
	})
}
