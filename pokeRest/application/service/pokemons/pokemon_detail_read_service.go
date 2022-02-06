package pokemons

import "github.com/Symthy/PokeRest/pokeRest/application/query"

type PokemonDetailReadService struct {
	qs query.IPokemonQueryService
}

// UC: 単体詳細取得
func FindPokemon(pokemonId uint) {
}
