package model

import (
	"weshierNext/handler"
	"weshierNext/pkg/auth"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/logger"
	"weshierNext/pkg/token"
	"weshierNext/pkg/validate"
	"weshierNext/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// UserModel user model
type UserModel struct {
	BaseModel
	UserName string `zh:"用户名" json:"userName" gorm:"not null;column:username;unique_index" binding:"required" validate:"min=3,max=10"`
	PassWord string `zh:"密码" json:"-" gorm:"not null;column:password;" binding:"required" validate:"min=5,max=50"`
	Email    string `zh:"邮箱" json:"email" gorm:"not null" binding:"required" validate:"email"`
	NickName string `zh:"昵称" json:"nickName" gorm:"column:nickname;" validate:"max=24"`
	Bio      string `zh:"简介" json:"bio"`
	Avatar   string `zh:"头像" json:"avatar"`
	URL      string `zh:"头像" json:"url"`
	Phone    uint64 `zh:"手机号" json:"phone"`
	Role     string `zh:"角色" json:"role"`
	Age      uint8  `zh:"年龄" json:"age"`
	Status   uint8  `json:"status"`
	Resume   uint8  `zh:"简历" json:"resume"`
	AuthID   uint64 `json:"authId" gorm:"column:auth_id;"`
}

// TableName specified table name
func (u *UserModel) TableName() string {
	return "ws_user"
}

// InsertAdminUser auto insert admin account into user table
func InsertAdminUser() (err error) {
	admin := &UserModel{
		UserName: viper.GetString("admin.userName"),
		PassWord: viper.GetString("admin.passWord"),
		Email:    viper.GetString("admin.email"),
		NickName: viper.GetString("admin.nickName"),
		Role:     viper.GetString("admin.role"),
	}
	if admin.NickName == "" {
		admin.NickName = util.RandomString(5)
	}
	err = admin.Validate()
	if err != nil {
		logger.Logger.Debug("admin user validate failed", zap.String("error", err.Error()))
		return
	}
	_, err = QueryUserByUsername(admin.UserName)
	if err == gorm.ErrRecordNotFound {
		admin.Create()
		return
	}
	if err != nil {
		logger.Logger.Panic("admin user query failed", zap.String("error", err.Error()))
	}
	return
}

// Create create a user
func (u *UserModel) Create() error {
	err := u.EncryptPwd()
	if err != nil {
		return err
	}
	return DB.Self.Create(&u).Error
}

// Delete delete a user
func (u *UserModel) Delete() error {
	return DB.Self.Delete(&u).Error
}

// DeleteByID delete a user by userID
func DeleteByID(id uint64) error {
	u := &UserModel{
		BaseModel: BaseModel{
			ID: id,
		},
	}
	return DB.Self.Delete(&u).Error
}

// QueryUserByUsername query user by username
func QueryUserByUsername(username string) (UserModel, error) {
	u := UserModel{}
	data := DB.Self.Where("username=?", username).First(&u)
	return u, data.Error
}

// QueryUser 查询用户
func QueryUser(condition string, args ...interface{}) (UserModel, error) {
	u := UserModel{}
	data := DB.Self.Where(condition, args...).First(&u)
	return u, data.Error
}

// QueryUserByUserID 根据 ID 查询用户信息
func QueryUserByUserID(id uint64) (UserModel, error) {
	u := UserModel{}
	data := DB.Self.Where("id=?", id).First(&u)
	return u, data.Error
}

// Compare compare pwd whether same
func (u *UserModel) Compare(pwd string) error {
	err := auth.Compare(u.PassWord, pwd)
	return err
}

// EncryptPwd encry user password
func (u *UserModel) EncryptPwd() (err error) {
	u.PassWord, err = auth.Encrypt(u.PassWord)
	return err
}

// Validate Validate the field
func (u *UserModel) Validate() error {
	return validate.Validate(*u)
}

// Login 登录操作
func (u *UserModel) Login(c *gin.Context) (claims *token.JWTClaims, t string, err error) {
	// sign the web token
	claims = &token.JWTClaims{
		ID:       u.ID,
		UserName: u.UserName,
		Email:    u.Email,
		NickName: u.NickName,
		Bio:      u.Bio,
		Avatar:   u.Avatar,
		URL:      u.URL,
		Phone:    u.Phone,
		Role:     u.Role,
		Age:      u.Age,
		Status:   u.Status,
		Resume:   u.Resume,
		AuthID:   u.AuthID,
	}
	t, err = token.Sign(c, *claims, viper.GetString("jwt.secret"))
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return claims, "", err
	}
	// save token to redis
	redisConn := DB.RedisPool.Get()
	defer redisConn.Close()
	// save token
	_, err = redisConn.Do("Set", t, t)
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return claims, "", err
	}
	// set expire time
	_, err = redisConn.Do("Expire", t, viper.GetInt64("jwt.maxage"))
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return claims, "", err
	}
	return claims, t, nil
}
