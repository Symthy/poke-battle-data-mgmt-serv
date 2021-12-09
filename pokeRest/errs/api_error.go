package errs

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiErrKey string

const (
	ApiErrAlreadyExistsUserName ApiErrKey = "ErrAlreadyExistsUserName"
	ApiErrUnexpected            ApiErrKey = "ErrUnexpected"
)

var (
	apiErrorMap = map[ApiErrKey]ApiError{
		ApiErrAlreadyExistsUserName: {status: http.StatusConflict, message: "指定したユーザ名は既に使用されています。別の名前を入力指定してください。", messageEN: "user name already exists"},

		ApiErrUnexpected: {status: http.StatusInternalServerError, message: "予期せぬエラーが発生しました。管理者に問い合わせてください。", messageEN: "unexpected error"},
	}
)

// visible for testing
func GetApiError(errKey ApiErrKey) ApiError {
	return apiErrorMap[errKey]
}

type PokeRestApiError interface {
	// APIError returns an HTTP status code and an API-safe error message.
	ApiError(errCode int) echo.HTTPError
}

type ApiError struct {
	status    int
	message   string
	messageEN string
}

func (e ApiError) ApiError(errCode string) echo.HTTPError {
	return echo.HTTPError{
		Code:    e.status,
		Message: e.buildResponseMessage(errCode),
	}
}

func (e ApiError) buildResponseMessage(errCode string) string {
	if errCode == "" {
		return e.message
	}
	return fmt.Sprintf("[%s] %s", errCode, e.message)
}
