package page

import (
	"weshierNext/handler"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Home 首页
func Home(c *gin.Context) {
	handler.SendResponse(c, errno.ErrNotFound, nil)
}

// NotFound 404 页面
func NotFound(c *gin.Context) {
	handler.SendResponse(c, errno.ErrNotFound, nil)
}
