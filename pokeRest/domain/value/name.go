package value

import (
	"regexp"

	"github.com/Symthy/PokeRest/pokeRest/errors"
)

type Name struct {
	value string
}

func NewName(value string) (*Name, error) {
	if len(value) > 16 || len(value) < 1 {
		return nil, errors.NewInvalidValueError("user name")
	}
	r := regexp.MustCompile(`[a-zA-Z0-9-_]+`)
	if !r.MatchString(value) {
		return nil, errors.NewInvalidValueError("user name")
	}
	return &Name{value: value}, nil
}

func (n Name) Value() string {
	return n.value
}
