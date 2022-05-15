package identifier

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/errs"
)

var _ entity.IValueId[uint64] = (*TrainedAdjustmentId)(nil)

type TrainedAdjustmentId struct {
	ValueId[uint64]
}

func NewTrainedAdjustmentId(value uint64) (*TrainedAdjustmentId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("TrainedAdjustmentId", "value", string(rune(value)))
	}
	return &TrainedAdjustmentId{ValueId[uint64]{value}}, nil
}
