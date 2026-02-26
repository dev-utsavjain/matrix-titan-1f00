package handlers

import "net/http"

// RegisterRoutes registers all generated API routes
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/posts", GetPosts)
	mux.HandleFunc("GET /api/posts/featured", GetFeaturedPosts)
	mux.HandleFunc("GET /api/posts/recent", GetRecentPosts)
	mux.HandleFunc("GET /api/posts/{slug}", GetPosts)
	mux.HandleFunc("GET /api/posts/related/{id}", GetRelatedPosts)
	mux.HandleFunc("GET /api/categories", GetCategories)
	mux.HandleFunc("GET /api/users/profile", GetUserProfile)
	mux.HandleFunc("GET /api/users/posts", GetPosts)
	mux.HandleFunc("GET /api/users/stats", GetUserStats)
}
