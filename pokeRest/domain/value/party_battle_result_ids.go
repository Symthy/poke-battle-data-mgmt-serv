package value

import "github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"

type PartyBattleResultIds struct {
	ids []identifier.PartyBattleResultId
}

func NewPartyBattleResultIds(values []uint) (*PartyBattleResultIds, error) {
	ids := make([]identifier.PartyBattleResultId, len(values))
	for i, v := range values {
		id, err := identifier.NewPartyBattleResultId(v)
		if err != nil {
			return nil, err
		}
		ids[i] = *id
	}
	return &PartyBattleResultIds{ids}, nil
}

func NewEmptyPartyBattleResultIds() PartyBattleResultIds {
	return PartyBattleResultIds{make([]identifier.PartyBattleResultId, 0)}
}
