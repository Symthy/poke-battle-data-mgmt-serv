package identifier

import "github.com/Symthy/PokeRest/pokeRest/errs"

type PokemonId struct {
	ValueId
}

func NewPokemonId(value uint) (*PokemonId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("PokemonId", "value", string(rune(value)))
	}
	return &PokemonId{ValueId{value}}, nil
}

func NewEmptyPokemonId() PokemonId {
	return PokemonId{ValueId{0}}
}
