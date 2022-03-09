package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type BattleRecordInput struct {
	id                            uint
	partyId                       uint
	generation                    int
	series                        int
	season                        int
	battleResult                  string
	selfElectionPokemonIds        []uint
	selfElectionTrainedPokemonIds []uint
	opponentElectionPokemonIds    []uint
	BattleOpponentPartyInput
}

func NewBattleRecordInput(
	id, partyId uint, generation, series, season int, battleResult string,
	selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds []uint,
	opponentPartyInput BattleOpponentPartyInput,
) BattleRecordInput {
	return BattleRecordInput{
		id:                            id,
		partyId:                       partyId,
		generation:                    generation,
		series:                        series,
		season:                        season,
		battleResult:                  battleResult,
		selfElectionPokemonIds:        selfElectionPokemonIds,
		selfElectionTrainedPokemonIds: selfElectionTrainedPokemonIds,
		opponentElectionPokemonIds:    opponentElectionPokemonIds,
		BattleOpponentPartyInput:      opponentPartyInput,
	}
}

func (b BattleRecordInput) BuildDomain() (*battles.BattleRecord, error) {
	id, err := identifier.NewBattleRecordId(b.id)
	if err != nil {
		return nil, err
	}
	partyId, err := identifier.NewPartyId(b.partyId)
	if err != nil {
		return nil, err
	}
	season, err := battles.NewSeason(b.generation, b.series, b.season)
	if err != nil {
		return nil, err
	}
	result := value.BattleResult(b.battleResult)
	selfPokemons := battles.NewElectionPokemons(b.selfElectionPokemonIds)
	selfTrainedPokemons := battles.NewElectionPokemons(b.selfElectionTrainedPokemonIds)
	opponentPokemons := battles.NewElectionPokemons(b.opponentElectionPokemonIds)
	opponentParty, err := b.BattleOpponentPartyInput.BuildDomain()
	if err != nil {
		return nil, err
	}
	domain := battles.NewBattleRecord(*id, *partyId, *season, result, selfPokemons, selfTrainedPokemons,
		opponentPokemons, *opponentParty)
	return &domain, nil
}
