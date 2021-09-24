package errno

/*
	错误码设计
	第一位表示错位分类, 1 为系统错误, 2 为普通错误
	第二三位表示服务模块代码
	第四五位表示具体错误代码
*/

// Errno 定义错误码
type Errno struct {
	Code    int
	Message string
}

func (err *Errno) Error() string {
	return err.Message
}

func (err Errno) ReplaceMsg(message string) Errno {
	err.Message = message
	return err
}

var (
	OK                  = &Errno{Code: 0, Message: "success"}
	InternalServerError = &Errno{Code: 10001, Message: "内部服务器错误"}
	ErrBind             = &Errno{Code: 10002, Message: "绑定请求体到 struct 时发生错误"}
	ErrRuntime          = &Errno{Code: 10003, Message: "程序运行错误"}

	UserSrvErr = &Errno{Code:    100101, Message: "用户服务错误"}

	ErrUserNotFound  = &Errno{Code: 20101, Message: "用户不存在"}
	ErrRequestParams = &Errno{Code: 20102, Message: "请求参数错误"}
	ErrUserIsExist  = &Errno{Code: 20103, Message: "用户已存在"}
	ErrUserPassword  = &Errno{Code: 20104, Message: "用户密码错误"}

)
