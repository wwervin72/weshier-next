package user

import (
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/token"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Login login
func Login(c *gin.Context) {
	var u LoginReqStruct
	if err := c.ShouldBindJSON(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	user, err := model.QueryUserByUsername(u.Username)
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	if len(user) == 0 {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	userInfo := user[0]
	if err = userInfo.Compare(u.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	// sign the web token
	t, err := token.Sign(c, token.JWTClaims{
		ID:       userInfo.ID,
		Username: userInfo.Username,
		Email:    userInfo.Email,
	}, viper.GetString("jwt.secret"))
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	// save token to redis
	redisConn := model.DB.RedisPool.Get()
	defer redisConn.Close()
	// save token
	_, err = redisConn.Do("Set", t, t)
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	// set expire time
	_, err = redisConn.Do("Expire", t, viper.GetInt64("jwt.maxage"))
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, &LoginResStruct{
		Token: t,
		UserModel: model.UserModel{
			Username: userInfo.Username,
			Email:    userInfo.Email,
			Nickname: userInfo.Nickname,
			BaseModel: model.BaseModel{
				ID: userInfo.ID,
			},
		},
	})
	return
}
