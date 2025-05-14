package codes

import "net/http"

const (
	InvalidParams               = uint32(40001)
	UserNoFound                 = uint32(40002)
	InvalidPassword             = uint32(40003)
	CompetitionQuestionNotFound = uint32(40004)
)

const (
	AccessTokenInvalid       = uint32(40101)
	RefreshTokenInvalid      = uint32(40102)
	CompetitionPasswordError = uint32(40103)
	AdminAuthFailed          = uint32(40104)
)

const (
	CompetitionNotRunning  = uint32(40301)
	CompetitionNotStarted  = uint32(40302)
	CompetitionNotAttend   = uint32(40303)
	BlockDuringCompetition = uint32(40304)
)

const (
	EmailAlreadyRegister     = uint32(40901)
	UserNameAlreadyRegister  = uint32(40902)
	CompetitionAlreadyAttend = uint32(40903)
)

const (
	ServerError      = uint32(50001)
	DBError          = uint32(50002)
	GenerateJWTError = uint32(50003)
	UnknownError     = uint32(50004)
)

var CodeToHTTPStatus = map[uint32]int{
	InvalidParams:               http.StatusBadRequest,
	UserNoFound:                 http.StatusBadRequest,
	InvalidPassword:             http.StatusBadRequest,
	CompetitionQuestionNotFound: http.StatusBadRequest,

	AccessTokenInvalid:       http.StatusUnauthorized,
	RefreshTokenInvalid:      http.StatusUnauthorized,
	CompetitionPasswordError: http.StatusUnauthorized,
	AdminAuthFailed:          http.StatusUnauthorized,

	BlockDuringCompetition: http.StatusForbidden,
	CompetitionNotAttend:   http.StatusForbidden,
	CompetitionNotStarted:  http.StatusForbidden,
	CompetitionNotRunning:  http.StatusForbidden,

	EmailAlreadyRegister:     http.StatusConflict,
	UserNameAlreadyRegister:  http.StatusConflict,
	CompetitionAlreadyAttend: http.StatusConflict,

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
