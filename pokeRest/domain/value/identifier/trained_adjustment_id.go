package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type TrainedAdjustmentId struct {
	ValueId
}

func NewTrainedAdjustmentId(value uint) (*TrainedAdjustmentId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("TrainedAdjustmentId", "value", string(rune(value)))
	}
	return &TrainedAdjustmentId{ValueId{value}}, nil
}
