package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type AbilityId struct {
	ValueId
}

func NewAbilityId(value uint) (*AbilityId, error) {
	if value < 0 {
		return nil, errs.ThrowErrorInvalidValue("AbilityId", "value", string(rune(value)))
	}
	return &AbilityId{ValueId{value}}, nil
}
