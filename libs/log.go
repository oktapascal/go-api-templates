package libs

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// CreateLogEntry creates a new logrus Entry with fields based on the provided HTTP request.
// If the request is nil, it will only include the current timestamp in the log entry.
// Otherwise, it will include the timestamp, HTTP method, request URI, and client IP address.
//
// Parameters:
// - r: A pointer to an http.Request object representing the incoming HTTP request.
//
// Returns:
// - A pointer to a logrus.Entry object containing the log entry fields.
func CreateLogEntry(r *http.Request) *logrus.Entry {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	if r == nil {
		return logger.WithFields(logrus.Fields{
			"at": time.Now().Format("2024-07-25 04:05:10"),
		})
	}

	return logger.WithFields(logrus.Fields{
		"at":     time.Now().Format("2024-07-25 04:05:10"),
		"method": r.Method,
		"uri":    r.RequestURI,
		"ip":     r.RemoteAddr,
	})
}
