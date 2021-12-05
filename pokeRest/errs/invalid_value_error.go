package errs

import "golang.org/x/xerrors"

type InvalidValueError struct {
	fieldName string
}

func NewInvalidValueError(fieldName string) error {
	return xerrors.Errorf("%w", &InvalidValueError{
		fieldName: fieldName,
	})
}

func (e InvalidValueError) Error() string {
	return "invalid value error: " + e.fieldName
}
