package errs

import "github.com/Symthy/PokeRest/internal/logs"

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
		ErrAuth:           initServerError(logs.Error, C1, "invalid token"),
		ErrAuthentication: initServerError(logs.Error, A1, "Authentication error. Invalid Password"),
		ErrUserNotFound:   initServerError(logs.Warn, A2, "user not found"),
		ErrInvalidValue:   initServerErrorWithFields(logs.Warn, D1, "invalid value.", "class,field,value"),
		ErrNoValue:        initServerErrorWithFields(logs.Error, D2, "no value", "class,field"),
		ErrUnexpected:     initServerError(logs.Error, C9999, "unexpected error"),
	}
)

func initServerError(level logs.Level, errCode ErrorCode, message string) ServerError {
	return initServerErrorWithFields(level, errCode, message, "")
}

func initServerErrorWithFields(level logs.Level, errCode ErrorCode, message string, fields string) ServerError {
	return ServerError{level: level, errCode: errCode, message: message, fields: fields}
}

// visible for testing
func GetServerError(errKey ServerErrKey) ServerError {
	return serverErrorMap[errKey]
}
