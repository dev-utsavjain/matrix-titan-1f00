package middleware

import (
	"backend/utils"
	"fmt"
	"log"
	"net/http"
)

func ErrorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				log.Printf("Request: %s %s", r.Method, r.URL.Path)
				
				utils.SendError(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
