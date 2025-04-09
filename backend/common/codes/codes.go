package codes

import "net/http"

const (
	InvalidParams   = uint32(40001)
	UserNoFound     = uint32(40002)
	InvalidPassword = uint32(40003)
)

const (
	JWTInvalid = uint32(40101)
	JWTExpired = uint32(40102)
)

const (
	EmailAlreadyRegister    = uint32(40901)
	UserNameAlreadyRegister = uint32(40902)
)

const (
	ServerError      = uint32(50001)
	DBError          = uint32(50002)
	GenerateJWTError = uint32(50003)
	UnknownError     = uint32(50004)
)

var CodeToHTTPStatus = map[uint32]int{
	InvalidParams:   http.StatusBadRequest,
	UserNoFound:     http.StatusBadRequest,
	InvalidPassword: http.StatusBadRequest,

	JWTInvalid: http.StatusUnauthorized,
	JWTExpired: http.StatusUnauthorized,

	EmailAlreadyRegister:    http.StatusConflict,
	UserNameAlreadyRegister: http.StatusConflict,

	ServerError:      http.StatusInternalServerError,
	DBError:          http.StatusInternalServerError,
	GenerateJWTError: http.StatusInternalServerError,
	UnknownError:     http.StatusInternalServerError,
}

func GetHTTPStatus(code uint32) int {
	if status, ok := CodeToHTTPStatus[code]; ok {
		return status
	}
	return http.StatusInternalServerError // 默认
}
