package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type TrainedDefenseTargetId struct {
	ValueId
}

func NewTrainedDefenseTargetId(value uint) (*TrainedDefenseTargetId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("TrainedDefenseTargetId", "value", string(rune(value)))
	}
	return &TrainedDefenseTargetId{ValueId{value}}, nil
}
