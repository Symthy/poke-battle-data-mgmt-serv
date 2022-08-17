package identifier

import (
	"github.com/Symthy/PokeRest/internal/domain/entity"
	"github.com/Symthy/PokeRest/internal/errs"
)

var _ entity.IValueId[uint64] = (*PartyId)(nil)

type PartyId struct {
	ValueId[uint64]
}

func NewPartyId(value uint64) (*PartyId, error) {
	if value < 0 {
		return nil, errs.ThrowErrorInvalidValue("PartyId", "value", string(rune(value)))
	}
	return &PartyId{ValueId[uint64]{value}}, nil
}

func NewEmptyPartyId() PartyId {
	return PartyId{}
}
