package utils

import (
	"github.com/gin-gonic/gin"
)

// SuccessResponse formats a successful API response
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	if data == nil {
		c.JSON(statusCode, gin.H{
			"message": message,
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

// ErrorResponse formats an error API response
func ErrorResponse(c *gin.Context, statusCode int, errorMsg string) {
	c.JSON(statusCode, gin.H{
		"error": errorMsg,
	})
}
