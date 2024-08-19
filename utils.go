package main

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func RespondWithError(c *gin.Context, logger *slog.Logger, statusCode int, message string, err error) {
	logger.Error(
		message,
		slog.String("error", err.Error()),
		slog.Int("status_code", statusCode),
	)
	c.JSON(statusCode, errorResponse{
		Message: message,
	})
}

func validateEmployee(e Employee) error {
	if e.Name == "" || e.Age < 0 || e.Salary < 0 {
		return fmt.Errorf("name and age must be valid")
	}
	return nil
}
