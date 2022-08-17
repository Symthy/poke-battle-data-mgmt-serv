package identifier

import (
	"github.com/Symthy/PokeRest/internal/domain/entity"
)

var _ entity.IValueId[uint64] = (*BattleRecordId)(nil)

type BattleRecordId struct {
	ValueId[uint64]
}

func NewBattleRecordId(value uint64) (*BattleRecordId, error) {
	// Todo: validate upper limit
	// if value < 0 {
	// 	return nil, errs.ThrowErrorInvalidValue("BattleRecordId", "value", string(rune(value)))
	// }
	return &BattleRecordId{ValueId[uint64]{value}}, nil
}
