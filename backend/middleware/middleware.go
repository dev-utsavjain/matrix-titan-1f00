package middleware

import (
	"net/http"
	"time"
)

func Chain(next http.Handler) http.Handler {
	return CORSMiddleware(
		LoggingMiddleware(
			ErrorHandlingMiddleware(next),
		),
	)
}
