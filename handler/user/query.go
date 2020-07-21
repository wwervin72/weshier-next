package user

import (
	"fmt"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/token"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// QueryUserInfo query userinfo by token
func QueryUserInfo(c *gin.Context) {
	u, err := token.ParseRequest(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	user, err := model.QueryUserByUserID(u.ID)
	if err == gorm.ErrRecordNotFound {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	} else if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	// 如果是非本地登录
	if user.AuthID != 0 {
		auth, err := model.QueryAuthByID(user.AuthID)
		if err == gorm.ErrRecordNotFound {
			handler.SendResponse(c, errno.ErrUserNotFound, nil)
			return
		} else if err != nil {
			handler.SendResponse(c, errno.InternalServerError, nil)
			return
		}
		fmt.Println(auth.AccessToken)
		// 根据 access_token 去 github 拉取用户信息
	}
	handler.SendResponse(c, nil, &user)
	return
}
