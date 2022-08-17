package identifier

import (
	"github.com/Symthy/PokeRest/internal/domain/entity"
	"github.com/Symthy/PokeRest/internal/errs"
)

var _ entity.IValueId[uint64] = (*PartyTagId)(nil)

type PartyTagId struct {
	ValueId[uint64]
}

func NewPartyTagId(value uint64) (*PartyTagId, error) {
	if value < 1 {
		return nil, errs.ThrowErrorInvalidValue("PartyTagId", "value", string(rune(value)))
	}
	return &PartyTagId{ValueId[uint64]{value}}, nil
}
