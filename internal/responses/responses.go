package responses

import (
	"github.com/NikitaKurabtsev/employee-api/internal/interfaces"
	"github.com/gin-gonic/gin"
)

type EmployeeResponse struct{}

func NewEmployeeResponse() *EmployeeResponse {
	return &EmployeeResponse{}
}

// RespondWithStatus sends a JSON response
// with the specified status code and data.
func (er EmployeeResponse) RespondWithStatus(
	c *gin.Context,
	statusCode int,
	data any,
) {
	c.JSON(statusCode, data)
}

// RespondWithError logs the error and sends a JSON response with
// the specified status code and message using Gin framework method.
func (er EmployeeResponse) RespondWithError(
	c *gin.Context,
	logger interfaces.Logger,
	statusCode int,
	customError string,
	originalErr error,
) {
	logger.Error(originalErr.Error(), "method", c.Request.Method)
	c.JSON(statusCode, gin.H{
		"error":   originalErr.Error(),
		"message": customError,
	})
}
