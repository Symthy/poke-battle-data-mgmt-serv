package battles

import (
	"time"

	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type IBattleRecordNotification interface {
	SetId(identifier.BattleRecordId)
	SetPartyId(identifier.PartyId)
	SetBattleResult(value.BattleResult)
	SetBattleOpponentPartyId(identifier.BattleOpponentPartyId)
	SetSelfElectionPokemons(ElectionPokemons[int])
	SetSelfTrainedPokemons(ElectionPokemons[uint])
	SetOpponentElectionPokemons(ElectionPokemons[int])
	IBattleSeasonNotification
}

type IBattleOpponentPartyNotification interface {
	SetId(identifier.BattleOpponentPartyId)
	SetOpponentPokemonIds(value.PartyPokemonIds)
}

type IBattleSeasonNotification interface {
	SetGeneration(int)
	SetSeries(int)
	SetSeason(int)
}

type IBattleSeasonPeriodNotification interface {
	IBattleSeasonNotification
	SetStartDateTime(time.Time)
	SetEndDateTime(time.Time)
}
