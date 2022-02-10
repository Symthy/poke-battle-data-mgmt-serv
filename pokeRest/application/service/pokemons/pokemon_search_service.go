package pokemons

import (
	"github.com/Symthy/PokeRest/pokeRest/application/query"
	"github.com/Symthy/PokeRest/pokeRest/application/service/pokemons/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
)

type PokemonMultiSearchService struct {
	qs query.IPokemonQueryService
}

// UC: 全検索
func (s PokemonMultiSearchService) SearchPokemons(cmd command.SearchPokemonCommand) (*pokemons.Pokemons, error) {
	return nil, nil
}
