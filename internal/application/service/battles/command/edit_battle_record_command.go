package command

import (
	"github.com/Symthy/PokeRest/internal/domain/factory"
)

type EditBattleRecordCommand struct {
	*factory.BattleRecordInput
}

func NewEditBattleRecordCommand(
	id, partyId, userId, generation, series, season uint64, battleResult string,
	selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds, opponentPartyPokemonIds []uint64,
) EditBattleRecordCommand {
	seasonInput := factory.NewSeasonBuilder().
		Generation(generation).Series(series).Season(season)

	opponentPartyInput := factory.NewBattleOpponentPartyBuilder().
		Id(id).OpponentPokemonIds(opponentPartyPokemonIds)

	return EditBattleRecordCommand{
		factory.NewBattleRecordBuilder().
			Id(id).
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
