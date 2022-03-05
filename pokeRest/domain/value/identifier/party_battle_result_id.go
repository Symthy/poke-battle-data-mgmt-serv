package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type PartyBattleResultId struct {
	ValueId
}

func NewPartyBattleResultId(value uint) (*PartyBattleResultId, error) {
	if value < 0 {
		return nil, errs.ThrowErrorInvalidValue("PartyBattleResultId", "value", string(rune(value)))
	}
	return &PartyBattleResultId{ValueId{value}}, nil
}
