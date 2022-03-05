package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

type EditBattleRecordCommand struct {
	factory.BattleRecordInput
}

func NewEditBattleRecordCommand(
	id uint, partyId uint, generation int, series int, season int, battleResult string,
	selfPokemonIds []int, selfTrainedPokemonIds []uint, opponentPokemonIds []int, opponentPartyMember []uint,
) EditBattleRecordCommand {
	return EditBattleRecordCommand{
		factory.NewBattleRecordInput(
			id, partyId, generation, series, season, battleResult,
			selfPokemonIds, selfTrainedPokemonIds, opponentPokemonIds, 0, opponentPartyMember),
	}
}
