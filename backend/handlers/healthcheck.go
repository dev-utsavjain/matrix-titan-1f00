package handlers

import (
	"net/http"
	"time"

	"backend/utils"
	"backend/views"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := views.HealthResponse{
		Success:   true,
		Message:   "Server is running",
		Timestamp: time.Now(),
	}
	utils.SendSuccess(w, response)
}
