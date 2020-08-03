package handler

import (
	"net/http"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/token"

	"github.com/gin-gonic/gin"
)

// Context custom gin context
type Context struct {
	*gin.Context
	User token.JWTClaims
}

// HandlerFunc handler fn
type HandlerFunc func(*Context)

// CustomContext wrapper gin context
func CustomContext(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context)
		context.Context = c
		if user, ok := c.Keys["user"]; ok {
			context.User = user.(token.JWTClaims)
		}
		handler(context)
	}
}

// Response 服务返回消息体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse 返回消息
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	// allways return http.StatusOK
	c.JSON(http.StatusOK, &Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
