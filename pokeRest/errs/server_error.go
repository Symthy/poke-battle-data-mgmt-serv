package errs

import (
	"fmt"

	pkgerrors "github.com/pkg/errors"
)

type ServerErrKey string

const (
	ErrAuth         ServerErrKey = "ErrAuth"
	ErrUserNotFound ServerErrKey = "ErrUserNotFound"
	ErrUnexpected   ServerErrKey = "ErrUnexpected"
)

var (
	serverErrorMap = map[ServerErrKey]ServerError{
		"ErrAuth":         {err: nil, level: Error, errCode: "0001", message: "invalid token"},
		"ErrUserNotFound": {err: nil, level: Warn, errCode: "0002", message: "user not found"},
		"ErrUnexpected":   {err: nil, level: Error, errCode: "9999", message: "unexpected error"},
	}
)

// visible for testing
func GetServerError(errKey ServerErrKey) ServerError {
	return serverErrorMap[errKey]
}

type ServerError struct {
	err        error
	level      Level
	errCode    string
	message    string
	stackTrace *string
}

func ThrowServerError(errKey ServerErrKey) error {
	e := GetServerError(errKey)
	if e.IsSaveOwnStackTrace() {
		trace := fmt.Sprintf("%+v", pkgerrors.New(""))
		e.stackTrace = &trace
	}
	return e
}

// dupplicate code. because don't increase the stack trace layer.
func WrapServerError(errKey ServerErrKey, target error) error {
	e := GetServerError(errKey)
	if e.IsSaveOwnStackTrace() {
		trace := fmt.Sprintf("%+v", pkgerrors.New(""))
		e.stackTrace = &trace
	}
	e.err = target
	return e
}

// func NewServerError(e error, level Level, errCode int, message string) *ServerError {
// 	return &ServerError{
// 		err:     e,
// 		level:   level,
// 		errCode: errCode,
// 		message: message,
// 	}
// }

// (private method) error interface implements
func (e ServerError) Error() string {
	var msg string
	if e.HasStackTrace() {
		msg = e.GetMessageAndStackTrace()
	} else {
		msg = e.GetMessage()
	}
	if e.err == nil {
		return msg
	}
	return msg + "\n" + e.err.Error()
}

func (e ServerError) Unwrap() error {
	return e.err
}

func (e ServerError) GetMessage() string {
	return fmt.Sprintf("[%5s - %04d] %s", e.level, e.errCode, e.message)
}

func (e ServerError) GetStackTrace() string {
	return *e.stackTrace
}

func (e ServerError) GetMessageAndStackTrace() string {
	return fmt.Sprintf("%s%s", e.GetMessage(), e.GetStackTrace())
}

func (e ServerError) GetErrorCode() string {
	return e.errCode
}

func (e ServerError) HasStackTrace() bool {
	return e.stackTrace != nil
}

func (e ServerError) IsSaveOwnStackTrace() bool {
	if e.level == Error || e.level == Fatal {
		return true
	}
	return false
}

func (e ServerError) IsNextError() bool {
	return e.err != nil
}

type IServerErrorAccessor interface {
	GetMessage() string
	GetStackTrace() string
	GetMessageAndStackTrace() string
	GetErrorCode() string
	HasStackTrace() bool
	IsSaveOwnStackTrace() bool
	IsNextError() bool
}

func AsServerError(target error) (IServerErrorAccessor, bool) {
	serverError, ok := target.(IServerErrorAccessor)
	return serverError, ok
}

func BuildErrorMessage(target error) string {
	serverError, isServerError := AsServerError(target)
	if !isServerError {
		return target.Error()
	}

	var errTopMsg string = ""
	var errDetailHeader string = ""
	var errDetailFooter string = ""
	errTopMsg = serverError.GetMessage()
	if serverError.IsNextError() {
		errDetailHeader = "\n---detail---\n"
		errDetailFooter = "\n------------\n"
	}
	errMsg := errTopMsg +
		errDetailHeader +
		target.Error() +
		errDetailFooter
	return errMsg
}
