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

// CreateLoggerConsole creates a logrus.Entry instance for logging to the console.
// It accepts an optional http.Request parameter, which is used to enrich the log output with request-specific information.
// If the request parameter is nil, the function logs a generic timestamp.
//
// The function initializes a logrus logger with the following settings:
// - Output: os.Stdout
// - Formatter: logrus.TextFormatter with ForceColors set to true
// - Level: logrus.TraceLevel
//
// The function returns a logrus.Entry instance with the following fields:
// - at: The current timestamp in the format "2024-07-25 04:05:10"
// - method: The HTTP method of the request (if the request parameter is not nil)
// - uri: The request URI (if the request parameter is not nil)
// - ip: The remote address of the request (if the request parameter is not nil)
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

// createLogsDir creates the logs directory if it does not exist.
//
// The function checks if the "logs" directory exists in the current working directory.
// If the directory does not exist, it attempts to create it with the specified permissions (0755).
//
// Parameters:
// None
//
// Return:
// error: An error if the directory creation fails, or nil if the directory already exists or was successfully created.
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

// CreateLoggerFile initializes and configures a logrus Logger instance for logging to a file.
// The function creates a logs directory if it does not exist, and then opens a log file with the current date as the filename.
// The Logger instance is configured with the following settings:
// - Formatter: logrus.JSONFormatter
// - Level: logrus.WarnLevel
// - Output: The opened log file
//
// If any errors occur during the directory creation or file opening process, the function logs a fatal error using the CreateLoggerConsole function.
//
// Parameters:
// None
//
// Return:
// *logrus.Logger: A configured logrus Logger instance for logging to a file
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
