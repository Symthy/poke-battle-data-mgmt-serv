package identifier

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/errs"
)

var _ entity.IValueId[uint64] = (*BattleOpponentPartyId)(nil)

type BattleOpponentPartyId struct {
	ValueId[uint64]
}

func NewBattleOpponentPartyId(value uint64) (*BattleOpponentPartyId, error) {
	if value < 0 {
		return nil, errs.ThrowErrorInvalidValue("BattleOpponentPartyId", "value", string(rune(value)))
	}
	return &BattleOpponentPartyId{ValueId[uint64]{value}}, nil
}

func NewEmptyBattleOpponentPartyId() BattleOpponentPartyId {
	return BattleOpponentPartyId{}
}
