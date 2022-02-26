package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type BattleOpponentPartyId struct {
	ValueId
}

func NewBattleOpponentPartyId(value uint) (*BattleOpponentPartyId, error) {
	if value < 0 {
		return nil, errs.ThrowServerErrorInvalidValue("BattleOpponentPartyId", "value", string(rune(value)))
	}
	return &BattleOpponentPartyId{ValueId{value}}, nil
}

func NewEmptyBattleOpponentPartyId() BattleOpponentPartyId {
	return BattleOpponentPartyId{}
}
