package user

import (
	"weshierNext/model"
)

// RegisterReqStruct register request body struct
type RegisterReqStruct struct {
	Username string `zh:"账号" json:"username" binding:"required" validate:"required,lt=5"`
	Password string `zh:"密码" json:"password" binding:"required" validate:"required,lt=5"`
}

// LoginReqStruct login request body struct
type LoginReqStruct struct {
	Username string `zh:"账号" json:"username" binding:"required"`
	Password string `zh:"密码" json:"password" binding:"required"`
}

// LoginResStruct login response body struct
type LoginResStruct struct {
	model.UserModel
	Token string `json:"token"`
}
