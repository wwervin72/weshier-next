package errno

import "fmt"

// Errno 自定义错误类型 通常用于返回给前端的错误
type Errno struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Err 自定义错误类型 通常用于日志打印错误
type Err struct {
	Code    int
	Message string
	Err     error
}

// Error 返回错误信息
func (err *Errno) Error() string {
	return err.Message
}

// New 创建 Err 实例
func New(errno *Errno, err error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Err:     err,
	}
}

// Add 添加错误信息
func (err *Err) Add(msg string) error {
	err.Message += " " + msg
	return err
}

// Addf 指定格式添加错误信息
func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

// Error 返回错误信息
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// IsErrUserNotFound 判断是否是用户找不到错误
func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

// DecodeErr 解析自定义的错误
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	}
	return InternalServerError.Code, err.Error()
}
