package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Login login
func Login(c *gin.Context) {
	var u LoginReqStruct
	if err := c.ShouldBindJSON(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	user, err := model.QueryUser("username=? and auth_id is NULL", u.Username)
	if err == gorm.ErrRecordNotFound {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	if err = user.Compare(u.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	t, err := user.Login(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, &LoginResStruct{
		Token: t,
		UserModel: model.UserModel{
			Username: user.Username,
			Email:    user.Email,
			Nickname: user.Nickname,
			BaseModel: model.BaseModel{
				ID: user.ID,
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
	client := &http.Client{}
	// 创建新的 http request，自定义 Header
	request, err := http.NewRequest("GET", fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s&redirect_uri=%s", viper.GetString("github.auth_url"),
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
	userReq, err := http.NewRequest("GET", fmt.Sprintf("%s?access_token=%s", viper.GetString("github.access_url"), data.AccessToken), nil)
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
	userExiested, err := model.QueryUser("username=? and auth_id is NOT NULL", userRespData.Username)
	var userAuth = &model.UserAuth{}
	if err == gorm.ErrRecordNotFound {
		// 如果是第一次登录
		// 需要在本地数据库创建一个对应账号
		userExiested = model.UserModel{
			Username: userRespData.Username,
			Email:    userRespData.Email,
			Bio:      userRespData.Bio,
			URL:      userRespData.URL,
			Avatar:   userRespData.Avatar,
			Nickname: userRespData.Name,
			Role:     model.TOURIST,
		}
		userAuth.OpenID = userRespData.ID
		userAuth.LoginType = "github"
		userAuth.AccessToken = data.AccessToken
		tx := model.DB.Self.Begin()
		insertUserResult := tx.Create(&userExiested)
		insertUserData, ok := insertUserResult.Value.(*model.UserModel)
		if insertUserResult.Error != nil || !ok {
			tx.Rollback()
			handler.SendResponse(c, errno.InternalServerError, nil)
			return
		}
		userAuth.UserID = insertUserData.ID
		insertUserAuthResult := tx.Create(&userAuth)
		insertUserAuthData, ok := insertUserAuthResult.Value.(*model.UserAuth)
		if insertUserAuthResult.Error != nil || !ok {
			tx.Rollback()
			handler.SendResponse(c, errno.InternalServerError, nil)
			return
		}
		userExiested.AuthID = insertUserAuthData.ID
		// 更新 authID 字段
		err = tx.Save(&userExiested).Error
		if err != nil {
			tx.Rollback()
			handler.SendResponse(c, errno.InternalServerError, nil)
			return
		}
		// 所有操作都完成
		tx.Commit()
	} else if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	userAuth, err = model.QueryAuthByID(userExiested.AuthID)
	// 更新 access_token
	userAuth.AccessToken = data.AccessToken
	err = model.DB.Self.Save(&userAuth).Error
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	t, err := userExiested.Login(c)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	handler.SendResponse(c, nil, &LoginResStruct{
		Token:     t,
		UserModel: userExiested,
		UserAuth:  *userAuth,
	})
}
