package identifier

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/errs"
)

var _ entity.IValueId[uint64] = (*TrainedPokemonId)(nil)

type TrainedPokemonId struct {
	ValueId[uint64]
}

func NewTrainedPokemonId(value uint64) (*TrainedPokemonId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("TrainedPokemonId", "value", string(rune(value)))
	}
	return &TrainedPokemonId{ValueId[uint64]{value}}, nil
}
