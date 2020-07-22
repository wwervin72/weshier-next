package model

// UserAuth 第三方登录关联表
type UserAuth struct {
	BaseModel
	// 本地数据库表 ws_user 存储的 Id
	UserID uint64 `json:"userId" gorm:"no null"`
	// 从第三方登录获取的用户信息里的 id
	OpenID    uint64 `json:"openId" gorm:""`
	LoginType string `json:"loginType" gorm:"no null"`
	// 第三方登录的 token
	AccessToken string `json:"access_token" gorm:"no null"`
}

// TableName 自定义表名称
func (ua *UserAuth) TableName() string {
	return "ws_user_auth"
}

// Create create a userAuth
func (ua *UserAuth) Create() error {
	return DB.Self.Create(&ua).Error
}

// QueryAuthByID 查询 auth
func QueryAuthByID(id uint64) (*UserAuth, error) {
	auth := &UserAuth{}
	data := DB.Self.Where("id=?", id).First(&auth)
	return auth, data.Error
}
