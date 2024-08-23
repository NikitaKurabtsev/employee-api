package httputils

import (
	"github.com/gin-gonic/gin"
)

func RespondWithStatus(
	c *gin.Context,
	statusCode int,
	data interface{},
) {
	c.JSON(statusCode, data)

}
