package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type MoveId struct {
	ValueId
}

func NewMoveId(value uint) (*MoveId, error) {
	if value < 0 {
		return nil, errs.ThrowErrorInvalidValue("MoveId", "value", string(rune(value)))
	}
	return &MoveId{ValueId{value}}, nil
}
