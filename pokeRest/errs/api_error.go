package errs

type RestApiError interface {
	// APIError returns an HTTP status code and an API-safe error message.
	ApiError() (int, string)
}

type ApiError struct {
	httpStatus  int
	responseMsg string
}
