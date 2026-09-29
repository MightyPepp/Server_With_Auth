package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("Запрос на: %s\n", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}