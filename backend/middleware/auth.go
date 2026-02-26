package middleware

import (
	"net/http"
	"strings"

	"backend/utils"
)

// AuthMiddleware validates authentication token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.SendError(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Extract token (Bearer token)
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.SendError(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := tokenParts[1]
		if token == "" {
			utils.SendError(w, "Token required", http.StatusUnauthorized)
			return
		}

		// TODO: Validate token and extract user ID
		// For now, we'll pass through with a placeholder
		// In production, validate JWT or session token here

		// Add user ID to request context
		// ctx := context.WithValue(r.Context(), "userID", userID)
		// next.ServeHTTP(w, r.WithContext(ctx))

		next.ServeHTTP(w, r)
	})
}
