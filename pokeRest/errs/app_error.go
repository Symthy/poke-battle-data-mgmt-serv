package errs

// import (
// 	"fmt"

// 	"github.com/pkg/errors"
// )

type AppError struct {
	ApiError
	ServerError
}

// // func NewAppError() AppError {
// // 	return
// // }

// // 一番下位層のメッセージから順に取り出す
// func (e AppError) Error() string {
// 	logMessage := ""
// 	if e.level == Error || e.level == Fatal {
// 		logMessage += e.getMessageAndStackTrace()
// 	} else {
// 		logMessage += e.getMessage()
// 	}

// 	nextAppError := AsAppError(e.next)
// 	if nextAppError != nil {
// 		return logMessage + nextAppError.Error()
// 	}

// 	builtInError := e.next
// 	if builtInError == nil {
// 		return logMessage
// 	}

// 	return logMessage + builtInError.Error()
// }

// func (e AppError) getMessageAndStackTrace() string {
// 	return e.getMessage() + e.getStackTrace()

// }

// func (e AppError) getStackTrace() string {
// 	return fmt.Sprintf("%+v\n", e)
// }

// func Wrap(err error, msg ...string) AppError {
// 	if err == nil {
// 		return nil
// 	}

// 	var m string
// 	if len(msg) != 0 {
// 		m = msg[0]
// 	}
// 	e := create(m)
// 	e.next = err
// 	return e
// }

// func Wrapf(err error, format string, args ...interface{}) AppError {
// 	e := create(fmt.Sprintf(format, args...))
// 	e.next = err
// 	return e
// }

// func As(err error, target interface{}) bool {
// 	return errors.As(err, target)
// }

// func AsAppError(err error) *AppError {
// 	if err == nil {
// 		return nil
// 	}

// 	var e *AppError
// 	if errors.As(err, &e) {
// 		return e
// 	}
// 	return nil
// }
