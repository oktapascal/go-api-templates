package middlewares

import (
	"go-rental/libs"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		libs.CreateLogEntry(r).Info("Incoming Request")
		next.ServeHTTP(w, r)
	})
}
