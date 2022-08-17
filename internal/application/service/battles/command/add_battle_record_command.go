package command

import (
	"github.com/Symthy/PokeRest/internal/domain/factory"
)

type AddBattleRecordCommand struct {
	*factory.BattleRecordInput
}

func NewAddBattleRecordOfCurrentCommand(
	partyId, userId uint64, battleResult string,
	selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds, opponentPartyPokemonIds []uint64,
) AddBattleRecordCommand {
	return NewAddBattleRecordCommand(
		partyId, userId, 0, 0, 0, battleResult, selfElectionPokemonIds, selfElectionTrainedPokemonIds,
		opponentElectionPokemonIds, opponentPartyPokemonIds)
}

func NewAddBattleRecordCommand(
	partyId, userId, generation, series, season uint64, battleResult string,
	selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds, opponentPartyPokemonIds []uint64,
) AddBattleRecordCommand {
	seasonInput := factory.NewSeasonBuilder().
		Generation(generation).Series(series).Season(season)
	opponentPartyInput := factory.NewBattleOpponentPartyBuilder().
		OpponentPokemonIds(opponentPartyPokemonIds)
	return AddBattleRecordCommand{
		factory.NewBattleRecordBuilder().
			PartyId(partyId).
			UserId(userId).
			BattleResult(battleResult).
			Season(seasonInput).
			SelfElectionPokemonIds(selfElectionPokemonIds).
			SelfElectionTrainedPokemonIds(selfElectionTrainedPokemonIds).
			OpponentElectionPokemonIds(opponentElectionPokemonIds).
			BattleOpponentParty(opponentPartyInput),
	}
}
