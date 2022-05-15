package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

type EditBattleRecordCommand struct {
	factory.BattleRecordInput
}

func NewEditBattleRecordCommand(
	id, partyId, userId, generation, series, season uint64, battleResult string,
	selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds, opponentPartyPokemonIds []uint64,
) EditBattleRecordCommand {
	return EditBattleRecordCommand{
		factory.NewBattleRecordInput(
			id, partyId, userId, generation, series, season, battleResult,
			selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds,
			factory.NewBattleOpponentPartyInput(0, opponentPartyPokemonIds...)),
	}
}
