package response

import "net/http"

type RespError struct {
	Code     int
	Msg      string
	HttpCode int
}

func NewRespError(code int, msg string, httpCode int) *RespError {
	return &RespError{
		Code:     code,
		Msg:      msg,
		HttpCode: httpCode,
	}
}

func (e *RespError) Error() string {
	return e.Msg
}

var (
	InvalidParams           = NewRespError(40001, "参数错误", http.StatusBadRequest)
	UserNoFound             = NewRespError(40002, "用户不存在", http.StatusBadRequest)
	InvalidPassword         = NewRespError(40003, "密码错误", http.StatusBadRequest)
	EmailAlreadyRegister    = NewRespError(40901, "邮箱已被注册", http.StatusConflict)
	UserNameAlreadyRegister = NewRespError(40902, "用户名已被注册", http.StatusConflict)
	ServerError             = NewRespError(50001, "服务器错误", http.StatusInternalServerError)
	DBError                 = NewRespError(50002, "数据库错误", http.StatusInternalServerError)
	GenerateJWTError        = NewRespError(50003, "JWT生成失败", http.StatusInternalServerError)
)
