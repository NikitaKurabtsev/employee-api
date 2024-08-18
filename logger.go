package main

import (
	"log/slog"
	"os"
)

func InitLogger() (*slog.Logger, error) {
	logFile, err := os.OpenFile("employee-api.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	logger := slog.New(slog.NewJSONHandler(logFile, nil))
	logger.Info("Application started")
	return logger, nil
}
