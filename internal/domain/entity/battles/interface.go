package battles

import (
	"time"

	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type IBattleRecordNotification interface {
	SetId(identifier.BattleRecordId)
	SetPartyId(identifier.PartyId)
	SetUserId(identifier.UserId)
	SetBattleResult(value.BattleResult)
	SetBattleOpponentPartyId(identifier.BattleOpponentPartyId)
	SetSelfElectionPokemons(*ElectionPokemons)
	SetSelfTrainedPokemons(*ElectionTrainedPokemons)
	SetOpponentElectionPokemons(*ElectionPokemons)
	IBattleOpponentPartyNotification
	IBattleSeasonNotification
}

type IBattleOpponentPartyNotification interface {
	SetOpponentPartyId(identifier.BattleOpponentPartyId)
	SetPokemonIds(*value.PartyPokemonIds)
}

type IBattleSeasonNotification interface {
	SetGeneration(uint16)
	SetSeries(uint16)
	SetSeason(uint16)
}

type IBattleSeasonPeriodNotification interface {
	IBattleSeasonNotification
	SetStartDateTime(time.Time)
	SetEndDateTime(time.Time)
}
