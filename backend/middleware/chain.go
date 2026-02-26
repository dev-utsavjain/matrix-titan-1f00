package middleware

import (
	"net/http"
)

func Chain(next http.Handler) http.Handler {
	return LoggingMiddleware(ErrorHandlingMiddleware(CORSMiddleware(next)))
}
