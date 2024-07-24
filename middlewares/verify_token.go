package middlewares

import (
	"go-rental/libs"
	"net/http"
	"strings"
)

func VerifyTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		token := strings.Replace(authorization, "Bearer", "", -1)
		token = strings.TrimSpace(token)

		err := libs.VerifyToken(token)

		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		next.ServeHTTP(w, r)
	})
}
