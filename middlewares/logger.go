package middlewares

import (
	"go-rental/config"
	"net/http"
)

// LoggerMiddleware is a middleware function that logs incoming HTTP requests.
// It wraps the provided http.Handler and logs a message using the CreateLogEntry function from the libs package.
//
// Parameters:
// - next: The http.Handler to be wrapped by the middleware.
//
// Returns:
// - An http.Handler that logs incoming requests before passing them to the next handler.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config.CreateLoggerConsole(r).Info("Incoming Request")
		next.ServeHTTP(w, r)
	})
}
