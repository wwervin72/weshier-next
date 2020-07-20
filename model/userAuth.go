package model

// UserAuth 第三方登录关联表
type UserAuth struct {
}

// TableName 自定义表名称
func (ua *UserAuth) TableName() string {
	return "ws_user_auth"
}

// Create create a userAuth
func (ua *UserAuth) Create() error {
	return DB.Self.Create(&ua).Error
}
