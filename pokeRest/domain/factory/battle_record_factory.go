package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/factory/inputs"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type BattleRecordFactory struct {
	input inputs.InputBattleRecord
}

func NewBattleRecordFactory(input inputs.InputBattleRecord) BattleRecordFactory {
	return BattleRecordFactory{input}
}

func (f BattleRecordFactory) CreateDomain() (*battles.BattleRecord, error) {
	id, err := value.NewBattleRecordId(f.input.Id())
	if err != nil {
		return nil, err
	}
	partyId, err := value.NewPartyId(f.input.PartyId())
	if err != nil {
		return nil, err
	}
	season := battles.NewSeason(f.input.Generation(), f.input.Series(), f.input.Season())
	result := value.BattleResult(f.input.BattleResult())
	selfPokemons := battles.NewElectionPokemons(f.input.SelfPokemonIds())
	selfTrainedPokemons := battles.NewElectionPokemons(f.input.SelfTrainedPokemonIds())
	opponentPokemons := battles.NewElectionPokemons(f.input.OpponentElectionPokemonIds())

	domain := battles.NewBattleRecord(*id, *partyId, season, result, selfPokemons, selfTrainedPokemons,
		value.NewEmptyBattleOpponentPartyId(), opponentPokemons)
	return &domain, nil
}
