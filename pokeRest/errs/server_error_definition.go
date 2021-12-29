package errs

import "github.com/Symthy/PokeRest/pokeRest/common"

type ServerErrKey string
type ErrorCode string

// Todo: out to file
const (
	ErrAuth           ServerErrKey = "ErrAuth"
	C1                ErrorCode    = "0001"
	ErrUnexpected     ServerErrKey = "ErrUnexpected"
	C9999             ErrorCode    = "9999"
	ErrAuthentication ServerErrKey = "ErrAuthentication"
	A1                ErrorCode    = "A0001"
	ErrUserNotFound   ServerErrKey = "ErrUserNotFound"
	A2                ErrorCode    = "A0002"
	ErrInvalidValue   ServerErrKey = "ErrInvalidValue"
	D1                ErrorCode    = "D0001"
	ErrNoValue        ServerErrKey = "ErrNoValue"
	D2                ErrorCode    = "D0002"
)

// Todo: out to file
var (
	serverErrorMap = map[ServerErrKey]ServerError{
		ErrAuth:           initServerError(common.Error, C1, "invalid token"),
		ErrAuthentication: initServerError(common.Error, A1, "Authentication error. Invalid Password"),
		ErrUserNotFound:   initServerError(common.Warn, A2, "user not found"), ErrInvalidValue: initServerErrorWithFields(common.Warn, D1, "invalid value.", "class,field,value"),
		ErrNoValue:    initServerErrorWithFields(common.Error, D2, "no value", "class,field"),
		ErrUnexpected: initServerError(common.Error, C9999, "unexpected error"),
	}
)

func initServerError(level common.Level, errCode ErrorCode, message string) ServerError {
	return initServerErrorWithFields(level, errCode, message, "")
}

func initServerErrorWithFields(level common.Level, errCode ErrorCode, message string, fields string) ServerError {
	return ServerError{level: level, errCode: errCode, message: message, fields: fields}
}

// visible for testing
func GetServerError(errKey ServerErrKey) ServerError {
	return serverErrorMap[errKey]
}
