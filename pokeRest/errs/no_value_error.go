package errs

import "golang.org/x/xerrors"

type NoValueError struct {
	fieldName string
}

func NewNoValueError(fieldName string) error {
	return xerrors.Errorf("%w", &NoValueError{
		fieldName: fieldName,
	})
}

func (e NoValueError) Error() string {
	return "No value error: " + e.fieldName
}
