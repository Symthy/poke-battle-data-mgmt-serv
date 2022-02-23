package command

import "github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"

type EditBattleRecordCommand struct {
	battles.BattleRecord
}

func NewEditBattleRecordCommand(
	id uint, partyId uint, battleFormat string, battleResult string, generation int, series int, season int,
	selfPokemonIds []int, selfTrainedPokemonIds []uint, opponentPokemonIds []int) AddBattleRecordCommand {
	return AddBattleRecordCommand{
		BattleRecord: battles.NewBattleRecord(
			id, partyId, generation, series, season, battleResult, 0,
			selfPokemonIds, selfTrainedPokemonIds, opponentPokemonIds),
	}
}

func (c EditBattleRecordCommand) ToDomain() battles.BattleRecord {
	return c.BattleRecord
}
