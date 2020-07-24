package errno

// Common errors
var (
	// OK 正常
	OK                  = &Errno{Code: 200, Message: ""}
	InternalServerError = &Errno{Code: 10001, Message: "服务器错误"}
	ErrBind             = &Errno{Code: 10002, Message: "参数错误."}
	ErrBadRequest       = &Errno{Code: 10003, Message: "错误的请求"}
	ErrNotFound         = &Errno{Code: 10004, Message: "资源未找到"}

	ErrValidation = &Errno{Code: 20001, Message: "校验失败"}
	ErrDatabase   = &Errno{Code: 20002, Message: "数据库错误"}
	ErrToken      = &Errno{Code: 20003, Message: "token 签名错误"}

	// user error
	ErrEncrypt           = &Errno{Code: 20101, Message: "密码加密失败"}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "用户未找到"}
	ErrTokenEmpty        = &Errno{Code: 20103, Message: "token 为空"}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "token 错误"}
	ErrTokenExpired      = &Errno{Code: 20104, Message: "token 已过期"}
	ErrPasswordIncorrect = &Errno{Code: 20105, Message: "密码不正确"}
)
