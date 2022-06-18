package pokemon

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

func Venusaur003() *pokemons.Pokemon {
	builder := factory.NewPokemonBuilder()
	builder.Id(3)
	builder.Name("フシギバナ")
	builder.EnglishName("Venusaur")
	builder.FormNo(1)
	builder.Generation(1)
	builder.BaseStatsH(80)
	builder.BaseStatsA(82)
	builder.BaseStatsB(83)
	builder.BaseStatsC(100)
	builder.BaseStatsD(100)
	builder.BaseStatsS(80)
	builder.AbilityIdOne(1)
	pokemon, _ := builder.BuildDomain()
	return pokemon
}
