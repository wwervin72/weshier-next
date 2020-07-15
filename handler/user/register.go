package user

import (
	"weshierNext/handler"

	"github.com/gin-gonic/gin"
)

// Register register user
func Register(c *gin.Context) {
	var body RegisterReqStruct
	var err error
	if err = c.ShouldBindJSON(&body); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
}
