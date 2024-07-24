package libs

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

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
