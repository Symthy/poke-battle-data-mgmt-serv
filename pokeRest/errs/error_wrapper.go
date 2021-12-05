package errs

type ErrorWrapper struct {
	next *ErrorWrapper
	err  error
}

func (ew ErrorWrapper) Message() string {
	message := ew.err.Error()
	if ew.next != nil {
		message += ew.next.Message()
	}
	return message
}
