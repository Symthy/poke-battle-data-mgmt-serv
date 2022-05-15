package value

import "github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"

type PartyTagIds struct {
	ids []identifier.PartyTagId
}

func NewPartyTagIds(values []uint64) (*PartyTagIds, error) {
	ids := make([]identifier.PartyTagId, len(values))
	for i, v := range values {
		id, err := identifier.NewPartyTagId(v)
		if err != nil {
			return nil, err
		}
		ids[i] = *id
	}
	return &PartyTagIds{ids}, nil
}
