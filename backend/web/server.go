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

func StartServer() {
	db.InitDB()

	server := createServer()

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
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	log.Printf("Server is ready to handle requests at %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
	}

	<-done
	log.Println("Server stopped")
}

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
