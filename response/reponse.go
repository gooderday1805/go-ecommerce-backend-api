package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {	
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Function handle response return json
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

// ErrorResponse handles error response return json
func ErrorResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
