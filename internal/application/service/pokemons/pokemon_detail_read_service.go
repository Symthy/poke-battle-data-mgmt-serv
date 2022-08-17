package pokemons

import "github.com/Symthy/PokeRest/internal/application/query"

type PokemonDetailReadService struct {
	qs query.IPokemonQueryService
}

// UC: 単体詳細取得
func FindPokemon(pokemonId uint) {
}
