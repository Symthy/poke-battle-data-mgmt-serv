package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

type AddBattleRecordCommand struct {
	factory.BattleRecordInput
}

func NewAddBattleRecordOfCurrentCommand(
	partyId uint, battleResult string, selfPokemonIds []uint, selfTrainedPokemonIds []uint,
	opponentPokemonIds []uint, opponentPartyMember []uint,
) AddBattleRecordCommand {
	return NewAddBattleRecordCommand(
		partyId, 0, 0, 0, battleResult, selfPokemonIds, selfTrainedPokemonIds,
		opponentPokemonIds, opponentPartyMember)
}

func NewAddBattleRecordCommand(
	partyId uint, generation int, series int, season int, battleResult string, selfPokemonIds []uint,
	selfTrainedPokemonIds []uint, opponentPokemonIds []uint, opponentPartyMember []uint,
) AddBattleRecordCommand {
	return AddBattleRecordCommand{
		factory.NewBattleRecordInput(
			0, partyId, generation, series, season, battleResult,
			selfPokemonIds, selfTrainedPokemonIds, opponentPokemonIds, 0, opponentPartyMember),
	}
}
