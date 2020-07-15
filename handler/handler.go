package handler

import (
	"net/http"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

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
