package errors

import (
	"github.com/pkg/errors"

	"golang.org/x/xerrors"
)

type UnauthorizedError struct {
	message string
}

func NewUnauthorizedError() UnauthorizedError {
	return errors.New("error")
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
