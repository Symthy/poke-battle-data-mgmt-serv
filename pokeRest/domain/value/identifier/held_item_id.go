package identifier

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/errs"
)

var _ entity.IValueId[uint16] = (*HeldItemId)(nil)

type HeldItemId struct {
	ValueId[uint16]
}

func NewHeldItemId(value uint64) (*HeldItemId, error) {
	// Todo: validate upper limit
	if value < 0 {
		return nil, errs.ThrowErrorInvalidValue("HeldItemId", "value", string(rune(value)))
	}
	return &HeldItemId{ValueId[uint16]{uint16(value)}}, nil
}
