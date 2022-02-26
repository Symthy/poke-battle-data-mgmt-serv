package identifier

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/errs"
)

var _ entity.IValueId = (*BattleRecordId)(nil)

type BattleRecordId struct {
	ValueId
}

func NewBattleRecordId(value uint) (*BattleRecordId, error) {
	if value < 0 {
		return nil, errs.ThrowServerErrorInvalidValue("BattleRecordId", "value", string(rune(value)))
	}
	return &BattleRecordId{ValueId{value}}, nil
}
