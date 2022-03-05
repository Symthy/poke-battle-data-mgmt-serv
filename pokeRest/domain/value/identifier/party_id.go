package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type PartyId struct {
	ValueId
}

func NewPartyId(value uint) (*PartyId, error) {
	if value < 0 {
		return nil, errs.ThrowErrorInvalidValue("PartyId", "value", string(rune(value)))
	}
	return &PartyId{ValueId{value}}, nil
}

func NewEmptyPartyId() PartyId {
	return PartyId{}
}
