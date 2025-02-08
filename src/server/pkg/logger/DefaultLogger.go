package logger

import (
	"io"
	"os"
	"path"

	"github.com/emperorsixpacks/dailbot/pkg/utils"
	"github.com/sirupsen/logrus"
)

var DefaultLogger *logrus.Logger

func NewDefaultLogger() {
	if DefaultLogger != nil {
		return
	}
	DefaultLogger = logrus.New()
	DefaultLogger.SetFormatter(&logrus.JSONFormatter{}) // JSON format for structured logs
	DefaultLogger.SetLevel(logrus.DebugLevel)           // Default log level

	// Log file setup
	logFile, err := utils.GetBasePath()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(path.Join(logFile, "logs/app.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		DefaultLogger.Warn("Failed to log to file, using default stderr")
		DefaultLogger.SetOutput(os.Stdout) // Output to console if file logging fails
	} else {
		multiWriter := io.MultiWriter(os.Stdout, file)
		DefaultLogger.SetOutput(multiWriter) // Output to both console and file
	}
}
