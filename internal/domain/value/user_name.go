package value

import (
	"regexp"

	"github.com/Symthy/PokeRest/internal/errs"
)

type UserName struct {
	value string
}

func NewUserName(value string) (*UserName, error) {
	if len(value) > 16 || len(value) < 1 {
		return nil, errs.ThrowErrorInvalidValue("User", "name", value)
	}
	r := regexp.MustCompile(`[a-zA-Z0-9-_]+`)
	if !r.MatchString(value) {
		return nil, errs.ThrowErrorInvalidValue("User", "name", value)
	}
	return &UserName{value: value}, nil
}

func (n UserName) Value() string {
	return n.value
}
