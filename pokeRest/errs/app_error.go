package errs

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

var (
	errMapping = map[string]ApiErrKey{
		"9999": ApiErrUnexpected,
	}
)

func resolveApiErrorKey(errCode string) ApiErrKey {
	return errMapping[errCode]
}

type IAppError interface {
	ApiErrorResponse() echo.HTTPError
	ServerErrorLogMsg() string
}

type AppError struct {
	serverError error
	errorCode   string
	apiError    ApiError
}

func newAppError(serverErr error, errCode string) AppError {
	return AppError{
		serverError: serverErr,
		errorCode:   errCode,
		apiError:    GetApiError(resolveApiErrorKey(errCode)),
	}
}

func newUnexpectedAppError(serverErr error) AppError {
	return AppError{
		serverError: serverErr,
		errorCode:   "",
		apiError:    GetApiError(ApiErrUnexpected),
	}
}

func BuildAppError(target error) AppError {
	serverErr, ok := AsServerError(target)
	if !ok {
		return newUnexpectedAppError(target)
	}
	return newAppError(target, serverErr.GetErrorCode())
}

func (e AppError) ApiErrorResponse() echo.HTTPError {
	status, message := e.apiError.ApiError()
	return echo.HTTPError{
		Code:    status,
		Message: e.buildErrorResponseMsg(e.errorCode, message),
	}
}

func (e AppError) buildErrorResponseMsg(errCode string, message string) string {
	if errCode == "" {
		return message
	}
	return fmt.Sprintf("[%s] %s", errCode, message)
}

func (e AppError) ServerErrorLogMsg() string {
	return e.serverError.Error()
}
