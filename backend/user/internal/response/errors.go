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
	EmailAlreadyRegister    = NewRespError(40001, "邮箱已被注册", http.StatusBadRequest)
	UserNameAlreadyRegister = NewRespError(40002, "用户名已被注册", http.StatusBadRequest)
	ServerError             = NewRespError(50001, "服务器错误", http.StatusInternalServerError)
	DBError                 = NewRespError(50002, "数据库错误", http.StatusInternalServerError)
)
