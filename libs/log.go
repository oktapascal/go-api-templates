package libs

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var loggerConsole = logrus.New()
var loggerFile = logrus.New()

func CreateLoggerConsole(r *http.Request) *logrus.Entry {
	loggerConsole.SetOutput(os.Stdout)
	loggerConsole.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	loggerConsole.SetLevel(logrus.TraceLevel)

	if r == nil {
		return loggerConsole.WithFields(logrus.Fields{
			"at": time.Now().Format("2024-07-25 04:05:10"),
		})
	}

	return loggerConsole.WithFields(logrus.Fields{
		"at":     time.Now().Format("2024-07-25 04:05:10"),
		"method": r.Method,
		"uri":    r.RequestURI,
		"ip":     r.RemoteAddr,
	})
}

func createLogsDir() error {
	logsDir := "logs"
	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		return os.Mkdir(logsDir, 0755)
	}
	return nil
}

func getLogFileName() string {
	return filepath.Join("logs", time.Now().Format("2006-01-02")+".log")
}

func CreateLoggerFile() *logrus.Logger {
	err := createLogsDir()
	if err != nil {
		CreateLoggerConsole(nil).Fatal("Failed to create logs directory")
	}

	logFileName := getLogFileName()

	file, errFile := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if errFile != nil {
		CreateLoggerConsole(nil).Fatal("Failed to open log file")
	}

	loggerFile.SetFormatter(&logrus.JSONFormatter{})
	loggerFile.SetLevel(logrus.WarnLevel)
	loggerFile.SetOutput(file)

	return loggerFile
}
