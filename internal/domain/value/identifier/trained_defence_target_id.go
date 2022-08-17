package identifier

import (
	"github.com/Symthy/PokeRest/internal/domain/entity"
	"github.com/Symthy/PokeRest/internal/errs"
)

var _ entity.IValueId[uint64] = (*TrainedDefenseTargetId)(nil)

type TrainedDefenseTargetId struct {
	ValueId[uint64]
}

func NewTrainedDefenseTargetId(value uint64) (*TrainedDefenseTargetId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("TrainedDefenseTargetId", "value", string(rune(value)))
	}
	return &TrainedDefenseTargetId{ValueId[uint64]{value}}, nil
}
