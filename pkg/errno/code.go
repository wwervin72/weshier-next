package errno

// Common errors
var (
	// OK 正常
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrBadRequest       = &Errno{Code: 10003, Message: "Bad request."}
	ErrNotFound         = &Errno{Code: 10004, Message: "Not found."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the json web token."}

	// user error
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user not found."}
	ErrTokenEmpty        = &Errno{Code: 20103, Message: "The token is empty."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token is invalid."}
	ErrTokenExpired      = &Errno{Code: 20104, Message: "The token is expired."}
	ErrPasswordIncorrect = &Errno{Code: 20105, Message: "The password was incorrect."}
)
