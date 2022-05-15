package pokemons

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type IPokemonNotification interface {
	SetId(identifier.PokemonId)
	SetPokedexId(value.PokedexId)
	SetFormName(string)
	SetName(string)
	SetEnglishName(string)
	SetGeneration(uint16)
	SetTypeSet(value.PokemonTypeSet)
	SetAbilitySet(value.PokemonAbilityIdSet)
	SetBaseStatsH(value.PokemonBaseStats)
	SetBaseStatsA(value.PokemonBaseStats)
	SetBaseStatsB(value.PokemonBaseStats)
	SetBaseStatsC(value.PokemonBaseStats)
	SetBaseStatsD(value.PokemonBaseStats)
	SetBaseStatsS(value.PokemonBaseStats)
	SetIsFinalEvolution(bool)
}
