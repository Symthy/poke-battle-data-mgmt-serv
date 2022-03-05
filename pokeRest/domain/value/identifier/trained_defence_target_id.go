package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type TrainedDefenceTargetId struct {
	ValueId
}

func NewTrainedDefenceTargetId(value uint) (*TrainedDefenceTargetId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("TrainedDefenceTargetId", "value", string(rune(value)))
	}
	return &TrainedDefenceTargetId{ValueId{value}}, nil
}
