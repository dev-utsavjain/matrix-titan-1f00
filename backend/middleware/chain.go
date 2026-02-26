package middleware

import "net/http"

// Chain applies multiple middleware functions in order
func Chain(handler http.Handler) http.Handler {
	// Apply middleware in reverse order (last applied first)
	handler = ErrorHandlingMiddleware(handler)
	handler = LoggingMiddleware(handler)
	handler = CORSMiddleware(handler)
	return handler
}
