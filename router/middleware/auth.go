package middleware

import (
	"github.com/gin-gonic/gin"
)

// IsLogin whether user is logined
func IsLogin(c *gin.Context) {
	// u, _, err := token.ParseRequest(c)
	// if err != nil {
	// 	handler.SendResponse(c, err, nil)
	// 	return
	// }

}

// IsAdmin whether user is admin role
func IsAdmin(c *gin.Context) {

}
