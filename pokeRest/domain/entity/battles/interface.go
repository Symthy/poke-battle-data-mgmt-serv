package battles

import (
	"time"

	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type IBattleRecordNotification interface {
	SetId(identifier.BattleRecordId)
	SetPartyId(identifier.PartyId)
	SetUserId(identifier.UserId)
	SetBattleResult(value.BattleResult)
	SetBattleOpponentPartyId(identifier.BattleOpponentPartyId)
	SetSelfElectionPokemons(ElectionPokemons)
	SetSelfTrainedPokemons(ElectionPokemons)
	SetOpponentElectionPokemons(ElectionPokemons)
	IBattleOpponentPartyNotification
	IBattleSeasonNotification
}

type IBattleOpponentPartyNotification interface {
	SetOpponentPartyId(identifier.BattleOpponentPartyId)
	SetPokemonIds(value.PartyPokemonIds)
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
