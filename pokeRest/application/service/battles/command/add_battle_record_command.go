package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type AddBattleRecordCommand struct {
	battles.BattleRecord
	opponentPartyMember value.PartyPokemonIds
}

func NewAddBattleRecordOfCurrentCommand(
	partyId uint, battleFormat string, battleResult string, selfPokemonIds []int, selfTrainedPokemonIds []uint,
	opponentPartyId uint, opponentPokemonIds []int, opponentElectionPokemonIds []int) AddBattleRecordCommand {
	return NewAddBattleRecordCommand(
		partyId, battleFormat, battleResult, 0, 0, 0, selfPokemonIds, selfTrainedPokemonIds,
		opponentPartyId, opponentPokemonIds, opponentElectionPokemonIds)
}

func NewAddBattleRecordCommand(
	partyId uint, battleFormat string, battleResult string, generation int, series int, season int,
	selfPokemonIds []int, selfTrainedPokemonIds []uint, opponentPartyId uint,
	opponentPokemonIds []int, opponentElectionPokemonIds []int) AddBattleRecordCommand {
	return AddBattleRecordCommand{
		BattleRecord: battles.NewBattleRecord(
			0, partyId, generation, series, season, battleResult, 0,
			selfPokemonIds, selfTrainedPokemonIds, opponentElectionPokemonIds),
		opponentPartyMember: value.NewPartyPokemonIds(opponentPokemonIds...),
	}
}

func (c AddBattleRecordCommand) ToDomain() battles.BattleRecord {
	return c.BattleRecord
}

func (c AddBattleRecordCommand) OpponentPartyMember() value.PartyPokemonIds {
	return c.opponentPartyMember
}
