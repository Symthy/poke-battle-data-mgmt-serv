package identifier

import (
	"github.com/Symthy/PokeRest/internal/domain/entity"
	"github.com/Symthy/PokeRest/internal/errs"
)

var _ entity.IValueId[uint16] = (*AbilityId)(nil)

type AbilityId struct {
	ValueId[uint16]
}

func NewAbilityId(value uint64) (*AbilityId, error) {
	if value < 0 {
		return nil, errs.ThrowErrorInvalidValue("AbilityId", "value", string(rune(value)))
	}
	return &AbilityId{ValueId[uint16]{uint16(value)}}, nil
}

func NewEmptyAbilityId() AbilityId {
	return AbilityId{newEmptyId[uint16]()}
}
