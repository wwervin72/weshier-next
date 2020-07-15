package user

import (
	"weshierNext/handler"
	"weshierNext/pkg/token"

	"github.com/gin-gonic/gin"
)

// QueryUserInfo query userinfo by token
func QueryUserInfo(c *gin.Context) {
	u, err := token.ParseRequest(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, &u)
	return
}
