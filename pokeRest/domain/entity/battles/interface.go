package battles

import "time"

type IBattleRecordNotification interface {
	Id(uint)
	PartyId(uint)
	Generation(int)
	Series(int)
	Season(int)
	BattleResult(string)
	BattleOpponentPartyId(uint)
	SelfElectionPokemons([]int)
	SelfTrainedPokemons([]uint)
	OpponentElectionPokemons([]int)
}

type IBattleSeasonNotification interface {
	Generation(int)
	Series(int)
	Season(int)
	StartDateTime(time.Time)
	EndDateTime(time.Time)
}
