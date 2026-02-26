package views

import "time"

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

type HealthResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
