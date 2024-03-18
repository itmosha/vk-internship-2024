package logger

import (
	"log"
	"log/slog"
	"net/http"
	"os"
)

type Logger struct {
	log *slog.Logger
}

// Setup slog logger.
func NewLogger(filePath, env string) *Logger {
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	var log *slog.Logger
	switch env {
	case "local":
		log = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return &Logger{log}
}

// Log an API request info.
func (l *Logger) Log(r *http.Request, status int, err error) {
	if status >= 500 { // Log server errors
		l.log.Error(
			"API request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", status),
			slog.String("error", err.Error()),
		)
	} else if status >= 400 { // Log client errors
		l.log.Info(
			"API request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", status),
			slog.String("error", err.Error()),
		)
	} else { // Log info
		l.log.Info(
			"API request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", status),
		)
	}
}
