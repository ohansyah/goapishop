package middleware

import (
	token "api_olshop/pkg/token"
	"log"
	"net/http"
)

// MiddlewareFunc standart mux
type MiddlewareFunc func(http.Handler) http.Handler

// LoggingMiddleware loging and validate token
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// log request
		log.Println(r.RequestURI)

		// token checks
		token.ValidateToken(w, r, next)
	})
}
