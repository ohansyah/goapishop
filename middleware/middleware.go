package middleware

import (
	token "api_olshop/token"
	"log"
	"net/http"
)

// MiddlewareFunc standart mux
type MiddlewareFunc func(http.Handler) http.Handler

// LoggingMiddleware loging and validate token
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)

		// insert log request api to db
		// insert res success request api to db
		// insert res failed request api to db

		// token checks
		token.ValidateToken(w, r, next)
	})
}
