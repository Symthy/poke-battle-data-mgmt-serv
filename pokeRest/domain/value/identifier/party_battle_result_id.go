package identifier

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/errs"
)

var _ entity.IValueId[uint64] = (*PartyBattleResultId)(nil)

type PartyBattleResultId struct {
	ValueId[uint64]
}

func NewPartyBattleResultId(value uint64) (*PartyBattleResultId, error) {
	if value < 0 {
		return nil, errs.ThrowErrorInvalidValue("PartyBattleResultId", "value", string(rune(value)))
	}
	return &PartyBattleResultId{ValueId[uint64]{value}}, nil
}
