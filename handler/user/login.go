package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

// GithubLogin github login
func GithubLogin(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	// 通过 code 在去请求 github 获取
	fmt.Println(code)
	client := &http.Client{}
	// 创建新的 http request，自定义 Header
	request, err := http.NewRequest("GET", fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
		viper.GetString("github.client_id"), viper.GetString("github.client_secret"), code, viper.GetString("github.redirect_url")), nil)
	// 设置 accept-type
	request.Header.Add("accept", "application/json")
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	resp, err := client.Do(request)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	// 注意关闭 body 防止内存泄漏
	defer resp.Body.Close()
	// 读取 res.body 中的数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	var data GithubAccessTokenRedirectStruct
	// 反序列化响应体内容，获取 access_token 内容
	err = json.Unmarshal(body, &data)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	// 根据 access_token 去获取用户信息
	// 创建请求
	userReq, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/user?access_token=%s", data.AccessToken), nil)
	userReq.Header.Add("accept", "application/json")
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	// 发送请求
	userResp, err := client.Do(userReq)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	defer userResp.Body.Close()
	// 读取响应体里的数据
	userRespBody, err := ioutil.ReadAll(userResp.Body)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	var userRespData GithubUserInfoStruct
	err = json.Unmarshal(userRespBody, &userRespData)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	userExiested, err := model.QueryUserByUsername(userRespData.Username)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	if len(userExiested) == 0 {
		// 如果是第一次登录
		// 需要在本地数据库创建一个对应账号
		// user := model.UserModel{
		// 	Username: userRespData.Username,
		// 	Password: ,
		// }
	} else {
		// 直接用该账号登录
	}
	handler.SendResponse(c, nil, userRespData)
}
