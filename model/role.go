package model

const (
	// ADMIN 管理员角色
	ADMIN = iota
	// MEMBER 普通账号角色
	MEMBER
	// TOURIST 游客角色
	TOURIST
)

// RoleModel user role
type RoleModel struct {
	RoleType uint8 `role`
}

// TableName 自定义表
func (rm *RoleModel) TableName() string {
	return "ws_role"
}

// Create create a role
func (rm *RoleModel) Create() error {
	return DB.Self.Create(&rm).Error
}
