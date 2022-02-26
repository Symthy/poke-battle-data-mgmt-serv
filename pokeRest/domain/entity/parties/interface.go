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
	SetPartyResultIds([]identifier.PartyBattleResultId)
	SetPartyTagIds([]identifier.PartyTagId)
	SetTrainedPokemons(value.PartyPokemonIds)
	SetUserId(identifier.UserId)
}

type IPartyBattleResultNotification interface {
	SetId(identifier.PartyBattleResultId)
	SetMaxRate(int)
	SetSeasonRanking(int)
	SetMaxSeasonRanking(int)
	SetWinCount(int)
	SetLoseCount(int)
	battles.IBattleSeasonNotification
}

type IPartyTagNotification interface {
	SetId(identifier.PartyTagId)
	SetName(string)
	SetIsGeneration(bool)
	SetIsSeason(bool)
}
