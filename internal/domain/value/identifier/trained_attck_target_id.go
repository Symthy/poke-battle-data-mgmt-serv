package identifier

import (
	"github.com/Symthy/PokeRest/internal/domain/entity"
	"github.com/Symthy/PokeRest/internal/errs"
)

var _ entity.IValueId[uint64] = (*TrainedAdjustmentId)(nil)

type TrainedAttackTargetId struct {
	ValueId[uint64]
}

func NewTrainedAttackTargetId(value uint64) (*TrainedAttackTargetId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("TrainedAttackTargetId", "value", string(rune(value)))
	}
	return &TrainedAttackTargetId{ValueId[uint64]{value}}, nil
}
