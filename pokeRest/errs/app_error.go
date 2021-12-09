package errs

var (
	errMapping = map[string]ApiErrKey{
		"9999": ApiErrUnexpected,
	}
)

func resolveApiErrorKey(errCode string) ApiErrKey {
	return errMapping[errCode]
}

type AppError struct {
	serverError error
	errorCode   string
	apiErr      ApiError
}

func newAppError(serverErr error, errCode string) AppError {
	return AppError{
		serverError: serverErr,
		errorCode:   errCode,
		apiErr:      GetApiError(resolveApiErrorKey(errCode)),
	}
}

func newUnexpectedAppError(serverErr error) AppError {
	return AppError{
		serverError: serverErr,
		errorCode:   "",
		apiErr:      GetApiError(ApiErrUnexpected),
	}
}

func ResolveAppError(target error) AppError {
	serverErr, ok := AsServerError(target)
	if !ok {
		return newUnexpectedAppError(target)
	}
	return newAppError(target, serverErr.GetErrorCode())
}
