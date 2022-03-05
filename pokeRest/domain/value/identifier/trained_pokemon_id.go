package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type TrainedPokemonId struct {
	ValueId
}

func NewTrainedPokemonId(value uint) (*TrainedPokemonId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("TrainedPokemonId", "value", string(rune(value)))
	}
	return &TrainedPokemonId{ValueId{value}}, nil
}
