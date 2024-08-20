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
	data interface{},
) {
	logger.Info(message, "data", data)
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

// OkMessage return formatted message for
// reponse with function name and message
func OkMessage(funcName, message string) string {
	okMessage := funcName + ": " + message
	return okMessage
}
