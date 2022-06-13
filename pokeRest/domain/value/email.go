package value

import (
	"regexp"

	"github.com/Symthy/PokeRest/pokeRest/errs"
)

type Email struct {
	value string
}

// Todo: bug
func NewEmail(value string) (*Email, error) {
	r := regexp.MustCompile(`^[a-zA-Z0-9_+-]+(.[a-zA-Z0-9_+-]+)*@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`)
	if !r.MatchString(value) {
		return nil, errs.ThrowErrorInvalidValue("User", "email", value)
	}
	return &Email{value: value}, nil
}

func (e Email) Value() string {
	return e.value
}
