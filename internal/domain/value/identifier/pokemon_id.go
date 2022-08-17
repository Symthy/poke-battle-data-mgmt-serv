package identifier

import (
	"github.com/Symthy/PokeRest/internal/domain/entity"
	"github.com/Symthy/PokeRest/internal/errs"
)

var _ entity.IValueId[uint16] = (*PokemonId)(nil)

type PokemonId struct {
	ValueId[uint16]
}

func NewPokemonId(value uint64) (*PokemonId, error) {
	// Todo: validate upper limit
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("PokemonId", "value", string(rune(value)))
	}
	return &PokemonId{ValueId[uint16]{uint16(value)}}, nil
}

func NewEmptyPokemonId() PokemonId {
	return PokemonId{ValueId[uint16]{0}}
}
