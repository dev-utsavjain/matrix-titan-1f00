package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"backend/db"
	"backend/handlers"
	"backend/middleware"
)

// spaHandler serves static files from dist/ and falls back to index.html for SPA routing
func spaHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/") {
		http.NotFound(w, r)
		return
	}

	// Try to serve static file from dist/
	distPath := "./dist"
	filePath := filepath.Join(distPath, r.URL.Path)

	// Check if file exists
	if info, err := os.Stat(filePath); err == nil && !info.IsDir() {
		http.ServeFile(w, r, filePath)
		return
	}

	// Fallback to index.html for SPA client-side routing
	http.ServeFile(w, r, filepath.Join(distPath, "index.html"))
}

// StartServer initializes and starts the HTTP server
func StartServer() {
	// Initialize dependencies
	db.InitDB()

	// Create server
	server := createServer()

	// Start with graceful shutdown
	runServer(server)
}

func createServer() *http.Server {
	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("GET /api/health", handlers.HealthCheck)

	// Register all generated API routes
	handlers.RegisterRoutes(mux)

	// SSR handler (serves dist/ and fallback to index.html)
	mux.HandleFunc("/", spaHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Addr:         ":" + port,
		Handler:      middleware.Chain(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func runServer(server *http.Server) {
	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
