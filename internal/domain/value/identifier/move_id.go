package identifier

import (
	"github.com/Symthy/PokeRest/internal/domain/entity"
	"github.com/Symthy/PokeRest/internal/errs"
)

var _ entity.IValueId[uint16] = (*MoveId)(nil)

type MoveId struct {
	ValueId[uint16]
}

func NewMoveId(value uint64) (*MoveId, error) {
	// Todo: validate upper limit
	if value < 0 {
		return nil, errs.ThrowErrorInvalidValue("MoveId", "value", string(rune(value)))
	}
	return &MoveId{ValueId[uint16]{uint16(value)}}, nil
}
