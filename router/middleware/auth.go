package middleware

import (
	"github.com/gin-gonic/gin"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/token"
)

// LoginRequired required user is logined
func LoginRequired(c *gin.Context) {
	u, _, err := token.ParseRequest(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	c.Set("user", u)
	c.Next()
}

// IsAdmin whether user is admin role
func IsAdmin(c *gin.Context) {
	u, ok := c.Get("suer")
	if !ok {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	user := u.(*token.JWTClaims)
	if user == nil {
		handler.SendResponse(c, errno.ErrTokenInvalid, nil)
		return
	}
	if user.Role != model.ADMIN {
		handler.SendResponse(c, errno.ErrNoPermission, nil)
		return
	}
	c.Next()
}
