package command

import "github.com/Symthy/PokeRest/pokeRest/domain/factory/inputs"

type AddBattleRecordCommand struct {
	inputs.InputBattleRecord
}

func NewAddBattleRecordOfCurrentCommand(
	partyId uint, battleResult string, selfPokemonIds []int, selfTrainedPokemonIds []uint,
	opponentPokemonIds []int, opponentPartyMember []int,
) AddBattleRecordCommand {
	return NewAddBattleRecordCommand(
		partyId, 0, 0, 0, battleResult, selfPokemonIds, selfTrainedPokemonIds,
		opponentPokemonIds, opponentPartyMember)
}

func NewAddBattleRecordCommand(
	partyId uint, generation int, series int, season int, battleResult string,
	selfPokemonIds []int, selfTrainedPokemonIds []uint, opponentPokemonIds []int, opponentPartyMember []int,
) AddBattleRecordCommand {
	return AddBattleRecordCommand{
		InputBattleRecord: inputs.NewInputBattleRecord(
			0, partyId, generation, series, season, battleResult,
			selfPokemonIds, selfTrainedPokemonIds, opponentPokemonIds, opponentPartyMember),
	}
}
