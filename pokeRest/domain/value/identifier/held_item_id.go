package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type HeldItemId struct {
	ValueId
}

func NewHeldItemId(value uint) (*HeldItemId, error) {
	if value < 0 {
		return nil, errs.ThrowErrorInvalidValue("HeldItemId", "value", string(rune(value)))
	}
	return &HeldItemId{ValueId{value}}, nil
}
