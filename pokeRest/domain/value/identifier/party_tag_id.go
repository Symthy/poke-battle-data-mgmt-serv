package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type PartyTagId struct {
	ValueId
}

func NewPartyTagId(value uint) (*PartyTagId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("PartyTagId", "value", string(rune(value)))
	}
	return &PartyTagId{ValueId{value}}, nil
}
