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
	InvalidParams               = NewError(codes.InvalidParams, "参数错误")
	UserNoFound                 = NewError(codes.UserNoFound, "用户不存在")
	InvalidPassword             = NewError(codes.InvalidPassword, "密码错误")
	CompetitionPasswordError    = NewError(codes.CompetitionPasswordError, "密码错误")
	CompetitionQuestionNotFound = NewError(codes.CompetitionQuestionNotFound, "题目不存在")
	OldPasswordError            = NewError(codes.OldPasswordError, "旧密码错误")
	NewPasswordError            = NewError(codes.NewPasswordError, "新密码格式不符合要求")
)

var (
	AccessTokenInvalid  = NewError(codes.AccessTokenInvalid, "AccessToken无效")
	RefreshTokenInvalid = NewError(codes.RefreshTokenInvalid, "RefreshToken无效")
)

var (
	CompetitionNotRunning  = NewError(codes.CompetitionNotRunning, "不在比赛时间段内")
	CompetitionNotStarted  = NewError(codes.CompetitionNotStarted, "比赛未开始")
	CompetitionNotAttend   = NewError(codes.CompetitionNotAttend, "比赛未报名")
	BlockDuringCompetition = NewError(codes.BlockDuringCompetition, "比赛中不能使用AI功能")
)

var (
	AdminAuthFailed = NewError(codes.AdminAuthFailed, "无权限")
)

var (
	EmailAlreadyRegister     = NewError(codes.EmailAlreadyRegister, "邮箱已被注册")
	UserNameAlreadyRegister  = NewError(codes.UserNameAlreadyRegister, "用户名已被注册")
	CompetitionAlreadyAttend = NewError(codes.CompetitionAlreadyAttend, "已经报过名了")
)

var (
	ServerError      = NewError(codes.ServerError, "服务器错误")
	DBError          = NewError(codes.DBError, "数据库错误")
	GenerateJWTError = NewError(codes.GenerateJWTError, "JWT生成失败")
	UnknownError     = NewError(codes.UnknownError, "未知错误")
)
