package identifier

import (
	"github.com/Symthy/PokeRest/internal/domain/entity"
)

var _ entity.IValueId[uint64] = (*PartyBattleResultId)(nil)

type PartyBattleResultId struct {
	ValueId[uint64]
}

func NewPartyBattleResultId(value uint64) (*PartyBattleResultId, error) {
	// Todo: validate upper limit
	// if value < 0 {
	// 	return nil, errs.ThrowErrorInvalidValue("PartyBattleResultId", "value", string(rune(value)))
	// }
	return &PartyBattleResultId{ValueId[uint64]{value}}, nil
}
