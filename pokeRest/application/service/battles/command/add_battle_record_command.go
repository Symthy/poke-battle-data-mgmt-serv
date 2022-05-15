package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

type AddBattleRecordCommand struct {
	factory.BattleRecordInput
}

func NewAddBattleRecordOfCurrentCommand(
	partyId, userId uint64, battleResult string,
	selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds, opponentPartyPokemonIds []uint64,
) AddBattleRecordCommand {
	return NewAddBattleRecordCommand(
		partyId, userId, 0, 0, 0, battleResult, selfElectionPokemonIds, selfElectionTrainedPokemonIds,
		opponentElectionPokemonIds, opponentPartyPokemonIds)
}

func NewAddBattleRecordCommand(
	partyId, userId, generation, series, season uint64, battleResult string,
	selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds, opponentPartyPokemonIds []uint64,
) AddBattleRecordCommand {
	return AddBattleRecordCommand{
		factory.NewBattleRecordInput(
			0, partyId, userId, generation, series, season, battleResult,
			selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds,
			factory.NewBattleOpponentPartyInput(0, opponentPartyPokemonIds...)),
	}
}
