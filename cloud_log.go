package bitmart

import (
	"fmt"
	"io"
	"log"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

type CustomLogger struct {
	logger   *log.Logger
	logLevel int
}

// NewCustomLogger Creating a custom logger
func NewCustomLogger(logLevel int, output io.Writer) *CustomLogger {
	// Create a new logger, setting the date, time and short file name as the default format
	return &CustomLogger{
		logger:   log.New(output, "", log.Ldate|log.Ltime),
		logLevel: logLevel,
	}
}

// SetLogLevel Setting the log level
func (l *CustomLogger) SetLogLevel(level int) {
	l.logLevel = level
}

// Log Print logs (support different log levels)
func (l *CustomLogger) Log(level int, msg string) {
	if level >= l.logLevel {
		switch level {
		case DEBUG:
			l.logger.Printf("[DEBUG] %s", msg)
		case INFO:
			l.logger.Printf("[INFO] %s", msg)
		case WARN:
			l.logger.Printf("[WARN] %s", msg)
		case ERROR:
			l.logger.Printf("[ERROR] %s", msg)
		default:
		}
	}
}

// Logf Create a log with formatting
func (l *CustomLogger) Logf(level int, format string, v ...interface{}) {
	if level >= l.logLevel {
		msg := fmt.Sprintf(format, v...)
		l.Log(level, msg)
	}
}
