package identifier

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/errs"
)

var _ entity.IValueId[uint64] = (*TrainedPokemonId)(nil)

type UserId struct {
	ValueId[uint64]
}

func NewUserId(value uint64) (*UserId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("UserId", "value", string(rune(value)))
	}
	return &UserId{ValueId[uint64]{value}}, nil
}

func NewEmptyUserId() UserId {
	return UserId{ValueId[uint64]{0}}
}
