package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

type EditBattleRecordCommand struct {
	factory.BattleRecordInput
}

func NewEditBattleRecordCommand(
	id uint, partyId uint, generation int, series int, season int, battleResult string,
	selfPokemonIds []uint, selfTrainedPokemonIds []uint, opponentPokemonIds []uint, opponentPartyMember []uint,
) EditBattleRecordCommand {
	return EditBattleRecordCommand{
		factory.NewBattleRecordInput(
			id, partyId, generation, series, season, battleResult,
			selfPokemonIds, selfTrainedPokemonIds, opponentPokemonIds, 0, opponentPartyMember),
	}
}
