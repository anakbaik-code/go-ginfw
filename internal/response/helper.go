package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, code int, message string, data any) {
	c.JSON(code, SuccessResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string, err any) {
	c.JSON(code, ErrorResponse{
		Status:  false,
		Message: message,
		Error:   err,
	})
}