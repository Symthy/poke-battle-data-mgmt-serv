package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type UserId struct {
	ValueId
}

func NewUserId(value uint) (*UserId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("UserId", "value", string(rune(value)))
	}
	return &UserId{ValueId{value}}, nil
}

func NewEmptyUserId() UserId {
	return UserId{ValueId{0}}
}
