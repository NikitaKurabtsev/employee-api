package errors

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

const (
	ErrEmployeeInvalidId  = "ID must be a number"
	ErrInvalidJSON        = "failed to parse JSON"
	ErrEmployeeUpdate     = "failed to update employee"
	ErrEmployeeNotFound   = "employee not found"
	ErrEmployeeValidation = "invalid employee data"
)

type errorResponse struct {
	Message string `json:"message"`
}

// RespondWithError logs the error and sends a JSON response
// with the specified status code and message using Gin framework method.
func RespondWithError(
	ginContext *gin.Context,
	logger *slog.Logger,
	statusCode int,
	message string,
	err error,
) {
	logger.Error(
		message,
		slog.String("error", err.Error()),
		slog.Int("status_code", statusCode),
	)
	ginContext.JSON(statusCode, errorResponse{
		Message: message,
	})
}

// ErrMessage return formatted message for
// reponse with function name and message
func ErrMessage(funcName, message string) string {
	errMessage := funcName + ": " + message
	return errMessage
}
