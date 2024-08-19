package errors

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidId      = "invalid employee ID"
	ErrInvalidJSON    = "failed to parse JSON"
	ErrUpdateEmployee = "failed to update employee"
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

func ErrMessage(funcName, message string) string {
	errMessage := funcName + message
	return errMessage
}
