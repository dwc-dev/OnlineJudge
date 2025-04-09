package errors

import "backend/common/codes"

type Error struct {
	Code uint32
	Msg  string
}

func NewError(code uint32, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (e *Error) Error() string {
	return e.Msg
}

var (
	InvalidParams   = NewError(codes.InvalidParams, "参数错误")
	UserNoFound     = NewError(codes.UserNoFound, "用户不存在")
	InvalidPassword = NewError(codes.InvalidPassword, "密码错误")
)

var (
	JWTInvalid = NewError(codes.JWTInvalid, "JWT无效")
	JWTExpired = NewError(codes.JWTExpired, "JWT过期")
)

var (
	EmailAlreadyRegister    = NewError(codes.EmailAlreadyRegister, "邮箱已被注册")
	UserNameAlreadyRegister = NewError(codes.UserNameAlreadyRegister, "用户名已被注册")
)

var (
	ServerError      = NewError(codes.ServerError, "服务器错误")
	DBError          = NewError(codes.DBError, "数据库错误")
	GenerateJWTError = NewError(codes.GenerateJWTError, "JWT生成失败")
	UnknownError     = NewError(codes.UnknownError, "未知错误")
)
