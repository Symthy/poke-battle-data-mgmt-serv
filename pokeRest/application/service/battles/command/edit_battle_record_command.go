package command

import "github.com/Symthy/PokeRest/pokeRest/domain/factory/inputs"

type EditBattleRecordCommand struct {
	inputs.InputBattleRecord
}

func NewEditBattleRecordCommand(
	id uint, partyId uint, generation int, series int, season int, battleResult string,
	selfPokemonIds []int, selfTrainedPokemonIds []uint, opponentPokemonIds []int, opponentPartyMember []int,
) EditBattleRecordCommand {
	return EditBattleRecordCommand{
		InputBattleRecord: inputs.NewInputBattleRecord(
			id, partyId, generation, series, season, battleResult,
			selfPokemonIds, selfTrainedPokemonIds, opponentPokemonIds, opponentPartyMember),
	}
}
