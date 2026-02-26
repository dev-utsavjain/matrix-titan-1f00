package handlers

import (
	"backend/utils"
	"net/http"
	"time"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message":   "Server is running",
		"timestamp": time.Now().Format(time.RFC3339),
	}
	utils.SendSuccess(w, response)
}