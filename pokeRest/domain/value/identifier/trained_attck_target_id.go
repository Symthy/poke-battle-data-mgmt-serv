package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type TrainedAttackTargetId struct {
	ValueId
}

func NewTrainedAttackTargetId(value uint) (*TrainedAttackTargetId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("TrainedAttackTargetId", "value", string(rune(value)))
	}
	return &TrainedAttackTargetId{ValueId{value}}, nil
}
