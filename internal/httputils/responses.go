package httputils

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func RespondWithStatus(
	c *gin.Context,
	logger *slog.Logger,
	statusCode int,
	message string,
	employeeId int,
) {
	logger.Info(message, "id", employeeId)
	c.JSON(statusCode, gin.H{
		"message": message,
		"id":      employeeId,
	})
}

// OkMessage return formatted message for
// reponse with function name and message
func OkMessage(funcName, message string) string {
	okMessage := funcName + ": " + message
	return okMessage
}
