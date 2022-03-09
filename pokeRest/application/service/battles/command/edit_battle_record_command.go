package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

type EditBattleRecordCommand struct {
	factory.BattleRecordInput
}

func NewEditBattleRecordCommand(
	id, partyId uint, generation, series, season int, battleResult string,
	selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds,
	opponentPartyPokemonIds []uint,
) EditBattleRecordCommand {
	return EditBattleRecordCommand{
		factory.NewBattleRecordInput(
			id, partyId, generation, series, season, battleResult,
			selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds,
			factory.NewBattleOpponentPartyInput(0, opponentPartyPokemonIds...)),
	}
}
