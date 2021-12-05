package errs

import (
	"golang.org/x/xerrors"
)

type UnauthorizedError struct {
	message string
}

func NewUnauthorizedError() UnauthorizedError {
	return UnauthorizedError{}
}
func (e UnauthorizedError) Error() string {
	return "An authentication error. Invalid Password"
}

type AuthorizedUserNameInvalidError struct {
}

func NewAuthorizedUserNameInvalidError() error {
	return xerrors.Errorf("%w", &UnauthorizedError{})
}
func (e AuthorizedUserNameInvalidError) Error() string {
	return "Invalid user name."
}
