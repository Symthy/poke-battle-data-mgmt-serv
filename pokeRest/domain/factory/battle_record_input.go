package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type BattleRecordInput struct {
	id                         uint
	partyId                    uint
	generation                 int
	series                     int
	season                     int
	battleResult               string
	selfPokemonIds             []uint
	selfTrainedPokemonIds      []uint
	opponentElectionPokemonIds []uint
	opponentPartyId            uint
	opponentPartyPokemonIds    []uint
}

func NewBattleRecordInput(
	id, partyId uint, generation, series, season int, battleResult string, selfPokemonIds []uint,
	selfTrainedPokemonIds []uint, opponentElectionPokemonIds []uint,
	opponentPartyId uint, opponentPartyPokemonIds []uint,
) BattleRecordInput {
	return BattleRecordInput{
		id:                         id,
		partyId:                    partyId,
		generation:                 generation,
		series:                     series,
		season:                     season,
		battleResult:               battleResult,
		selfPokemonIds:             selfPokemonIds,
		selfTrainedPokemonIds:      selfTrainedPokemonIds,
		opponentElectionPokemonIds: opponentElectionPokemonIds,
		opponentPartyId:            opponentPartyId,
		opponentPartyPokemonIds:    opponentPartyPokemonIds,
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
	season := battles.NewSeason(b.generation, b.series, b.season)
	result := value.BattleResult(b.battleResult)
	selfPokemons := battles.NewElectionPokemons(b.selfPokemonIds)
	selfTrainedPokemons := battles.NewElectionPokemons(b.selfTrainedPokemonIds)
	opponentPokemons := battles.NewElectionPokemons(b.opponentElectionPokemonIds)
	var opponentPartyId identifier.BattleOpponentPartyId
	if b.opponentPartyId > 0 {
		oppId, err := identifier.NewBattleOpponentPartyId(b.opponentPartyId)
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

func (i BattleRecordInput) OpponentPartyPokemonIds() []uint {
	return i.opponentPartyPokemonIds
}
