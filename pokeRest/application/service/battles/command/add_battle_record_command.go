package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

type AddBattleRecordCommand struct {
	factory.BattleRecordInput
}

func NewAddBattleRecordOfCurrentCommand(
	partyId uint, battleResult string, selfElectionPokemonIds, selfElectionTrainedPokemonIds,
	opponentElectionPokemonIds, opponentPartyPokemonIds []uint,
) AddBattleRecordCommand {
	return NewAddBattleRecordCommand(
		partyId, 0, 0, 0, battleResult, selfElectionPokemonIds, selfElectionTrainedPokemonIds,
		opponentElectionPokemonIds, opponentPartyPokemonIds)
}

func NewAddBattleRecordCommand(
	partyId uint, generation, series, season int, battleResult string,
	selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds,
	opponentPartyPokemonIds []uint,
) AddBattleRecordCommand {
	return AddBattleRecordCommand{
		factory.NewBattleRecordInput(
			0, partyId, generation, series, season, battleResult,
			selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds,
			factory.NewBattleOpponentPartyInput(0, opponentPartyPokemonIds...)),
	}
}
