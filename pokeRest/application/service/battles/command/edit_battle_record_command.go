package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type EditBattleRecordCommand struct {
	AddBattleRecordCommand
}

func NewEditBattleRecordCommand(
	id uint, partyId uint, battleFormat string, battleResult string, generation int, series int, season int,
	selfPokemonIds []int, selfTrainedPokemonIds []uint, opponentPokemonIds []int) EditBattleRecordCommand {
	return EditBattleRecordCommand{
		AddBattleRecordCommand{
			battleRecord: battles.NewBattleRecord(
				id, partyId, generation, series, season, battleResult, 0,
				selfPokemonIds, selfTrainedPokemonIds, opponentPokemonIds),
			opponentPartyMember: value.NewPartyPokemonIds(opponentPokemonIds...),
		},
	}
}
