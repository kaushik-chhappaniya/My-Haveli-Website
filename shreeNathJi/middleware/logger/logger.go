package logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Log levels
const (
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
	FATAL   = "FATAL"
)

func Logger(str string) {
	logWithLevel(DEBUG, str)
}

// Infof logs an info message with formatting
func Infof(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	logWithLevel(INFO, message)
}

// Info logs an info message
func Info(str string) {
	logWithLevel(INFO, str)
}

// Warnf logs a warning message with formatting
func Warnf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	logWithLevel(WARNING, message)
}

// Warn logs a warning message
func Warn(str string) {
	logWithLevel(WARNING, str)
}

// Errorf logs an error message with formatting
func Errorf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	logWithLevel(ERROR, message)
}

// Error logs an error message
func Error(str string) {
	logWithLevel(ERROR, str)
}

// Fatalf logs a fatal message with formatting and exits
func Fatalf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	logWithLevel(FATAL, message)
	os.Exit(1)
}

// Fatal logs a fatal message and exits
func Fatal(str string) {
	logWithLevel(FATAL, str)
	os.Exit(1)
}

// Debugf logs a debug message with formatting
func Debugf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	logWithLevel(DEBUG, message)
}

// Debug logs a debug message
func Debug(str string) {
	logWithLevel(DEBUG, str)
}

func logWithLevel(level, str string) {
	verbosity := strings.ToLower(os.Getenv("LOG_VERBOSE"))

	switch verbosity {
	case "debug", "console", "1":
		// Console debug logs only
		log.Printf("%s LOGS: %s", level, str)

	case "file", "2":
		// File logs only
		writeToFile(level, str)

	case "both", "all", "3":
		// Both console and file logs
		log.Printf("%s LOGS: %s", level, str)
		writeToFile(level, str)

	case "true": // Backward compatibility
		log.Printf("%s LOGS: %s", level, str)

	case "false", "0", "":
		// Do nothing

	default:
		// Default to console logging
		log.Printf("%s LOGS: %s", level, str)
	}
}

func writeToFile(level, message string) {
	// Create logs directory if it doesn't exist
	logsDir := "logs"
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		log.Printf("Error creating logs directory: %v", err)
		return
	}

	// Create filename with current date
	currentTime := time.Now()
	filename := fmt.Sprintf("app-%s.log", currentTime.Format("2006-01-02"))
	filePath := filepath.Join(logsDir, filename)

	// Open file in append mode
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer file.Close()

	// Write log with timestamp and level
	logEntry := fmt.Sprintf("[%s] %s LOGS: %s\n",
		currentTime.Format("2006-01-02 15:04:05"), level, message)

	if _, err := file.WriteString(logEntry); err != nil {
		log.Printf("Error writing to log file: %v", err)
	}
}

// LoggingMiddleware logs HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Create a custom response writer to capture status code
		wrappedWriter := &responseWriter{
			ResponseWriter: w,
			statusCode:     0, // Initialize as 0, will be set when WriteHeader is called
			headerWritten:  false,
		}

		// Log request details
		logMessage := fmt.Sprintf("Request: %s %s from %s",
			r.Method, r.URL.Path, r.RemoteAddr)
		Info(logMessage)

		// Call the next handler
		next.ServeHTTP(wrappedWriter, r)

		// If no status was set, assume 200
		status := wrappedWriter.statusCode
		if status == 0 {
			status = http.StatusOK
		}

		// Log response details
		duration := time.Since(startTime)
		responseLog := fmt.Sprintf("Response: %s %s | Status: %d | Duration: %v",
			r.Method, r.URL.Path, status, duration)

		// Log with appropriate level based on status code
		if status >= 500 {
			Error(responseLog)
		} else if status >= 400 {
			Warn(responseLog)
		} else {
			Info(responseLog)
		}
	})
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode    int
	headerWritten bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if !rw.headerWritten {
		rw.statusCode = code
		rw.ResponseWriter.WriteHeader(code)
		rw.headerWritten = true
	}
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	// If WriteHeader hasn't been called yet, it will be called with 200
	if !rw.headerWritten {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(b)
}
