package middlewares

import (
	"net/http"
)

func AuthorizationCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		// Check if the request has a valid authorization token
		// If not, return a 401 Unauthorized response
		if authorization == "" {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		next.ServeHTTP(w, r)
	})
}
