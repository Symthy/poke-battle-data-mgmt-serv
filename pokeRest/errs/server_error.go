package errs

import (
	"fmt"

	pkgerrors "github.com/pkg/errors"
)

var (
	ErrAuth         = "ErrAuth"
	ErrUserNotFound = "ErrUserNotFound"
	ErrUnexpected   = "ErrUnexpected"

	errorList = map[string]*ServerError{
		"ErrAuth":         {err: nil, level: Error, errCode: 1, message: "invalid token"},
		"ErrUserNotFound": {err: nil, level: Warn, errCode: 2, message: "user not found"},
		"ErrUnexpected":   {err: nil, level: Error, errCode: 9999, message: "unexpected error"},
	}
)

// visible for testing
func GetServerError(errKey string) *ServerError {
	return errorList[errKey]
}

type ServerError struct {
	err        error
	level      Level
	errCode    int
	message    string
	stackTrace *string
}

type errorMessageHolder interface {
	getMessage() string
	isNextError() bool
}

func ThrowServerError(errKey string) error {
	e := GetServerError(errKey)
	if e.isSaveOwnStackTrace() {
		trace := fmt.Sprintf("%+v", pkgerrors.New(""))
		e.stackTrace = &trace
	}
	return e
}

// dupplicate code. because don't increase the stack trace layer.
func WrapServerError(errKey string, target error) error {
	e := GetServerError(errKey)
	if e.isSaveOwnStackTrace() {
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
	if e.stackTrace == nil {
		msg = e.getMessage()
	} else {
		msg = e.getMessageAndStackTrace()
	}
	return msg + "\n" + e.err.Error()
}

func BuildErrorMessage(target error) string {
	var errTopMsg string = ""
	var errDetailHeader string = ""
	var errDetailFooter string = ""
	if errMsgHolder, ok := target.(errorMessageHolder); ok {
		errTopMsg = errMsgHolder.getMessage()
		if errMsgHolder.isNextError() {
			errDetailHeader = "\n---detail---\n"
			errDetailFooter = "\n------------\n"
		}
	}

	errMsg := errTopMsg +
		errDetailHeader +
		target.Error() +
		errDetailFooter
	return errMsg
}

func (e ServerError) getMessage() string {
	return fmt.Sprintf("[%5s - %04d] %s", e.level, e.errCode, e.message)
}

func (e ServerError) getMessageAndStackTrace() string {
	return fmt.Sprintf("%s%s", e.getMessage(), *e.stackTrace)
}

func (e ServerError) isSaveOwnStackTrace() bool {
	if e.level == Error || e.level == Fatal {
		return true
	}
	return false
}

func (e ServerError) isNextError() bool {
	return e.err != nil
}
