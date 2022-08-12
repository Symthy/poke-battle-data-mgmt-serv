package parties

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type IPartyNotification interface {
	SetId(identifier.PartyId)
	SetName(string)
	SetBattleFormat(value.BattleFormat)
	SetIsPrivate(bool)
	// SetPartyResultIds(*value.PartyBattleResultIds)
	// SetPartyTagIds(*value.PartyTagIds)
	// SetTrainedPokemons(*value.PartyPokemonIds)
	// SetUserId(identifier.UserId)
}

type IPartyBattleResultNotification interface {
	SetId(identifier.PartyBattleResultId)
	SetMaxRate(uint16)
	SetSeasonRanking(uint16)
	SetMaxSeasonRanking(uint16)
	SetWinCount(uint16)
	SetLoseCount(uint16)
	battles.IBattleSeasonNotification
}

type IPartyTagNotification interface {
	SetId(identifier.PartyTagId)
	SetName(string)
	SetIsGeneration(bool)
	SetIsSeason(bool)
}
