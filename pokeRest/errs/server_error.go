package errs

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrAuth         = &ServerError{level: Error, errCode: 1, message: "invalid token"}
	ErrUserNotFound = &ServerError{level: Warn, errCode: 1, message: "user not found"}
)

type ServerError struct {
	error
	level   Level
	errCode int
	message string
}

func NewServerError(err error, level Level, errCode int, message string) *ServerError {
	return &ServerError{
		error:   err,
		level:   level,
		errCode: errCode,
		message: message,
	}
}

func (e ServerError) Error() string {
	if e.error == nil {
		return e.getMessage()
	}

	return fmt.Sprintf("%s\n%+v", e.getMessage(), e.error)
}

func (e *ServerError) Wrap(target error) {

	if e.error == nil {
		e.error = target
		return
	}
	e.error = errors.Wrap(e.error, target.Error())
}

func (e ServerError) getMessage() string {
	return fmt.Sprintf("[%s - %04d] %s\n", e.level, e.errCode, e.message)
}

func (e ServerError) isSaveOwnStackTrace(target error) bool {
	var serverError ServerError
	if errors.As(target, &serverError) {
		if serverError.level == Error || serverError.level == Fatal {
			return true
		}
		return false
	}
	return true
}
