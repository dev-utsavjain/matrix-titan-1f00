package middleware

import "net/http"

func Chain(next http.Handler) http.Handler {
	return CORSMiddleware(
		ErrorHandlingMiddleware(
			LoggingMiddleware(next),
		),
	)
}
