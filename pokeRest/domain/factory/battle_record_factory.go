package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/factory/inputs"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type BattleRecordFactory struct {
	input inputs.InputBattleRecord
}

func NewBattleRecordFactory(input inputs.InputBattleRecord) BattleRecordFactory {
	return BattleRecordFactory{input}
}

func (f BattleRecordFactory) CreateDomain() (*battles.BattleRecord, error) {
	id, err := identifier.NewBattleRecordId(f.input.Id())
	if err != nil {
		return nil, err
	}
	partyId, err := identifier.NewPartyId(f.input.PartyId())
	if err != nil {
		return nil, err
	}
	season := battles.NewSeason(f.input.Generation(), f.input.Series(), f.input.Season())
	result := value.BattleResult(f.input.BattleResult())
	selfPokemons := battles.NewElectionPokemons(f.input.SelfPokemonIds())
	selfTrainedPokemons := battles.NewElectionPokemons(f.input.SelfTrainedPokemonIds())
	opponentPokemons := battles.NewElectionPokemons(f.input.OpponentElectionPokemonIds())
	var opponentPartyId identifier.BattleOpponentPartyId
	if f.input.OpponentPartyId() > 0 {
		oppId, err := identifier.NewBattleOpponentPartyId(f.input.OpponentPartyId())
		if err != nil {
			return nil, err
		}
		opponentPartyId = *oppId
	} else {
		opponentPartyId = identifier.NewEmptyBattleOpponentPartyId()
	}

	domain := battles.NewBattleRecord(*id, *partyId, season, result, selfPokemons, selfTrainedPokemons,
		opponentPartyId, opponentPokemons)
	return &domain, nil
}
