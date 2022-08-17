package pokemons

import (
	"github.com/Symthy/PokeRest/internal/application/query"
	"github.com/Symthy/PokeRest/internal/application/service/pokemons/command"
	"github.com/Symthy/PokeRest/internal/domain/entity/pokemons"
)

type PokemonMultiSearchService struct {
	qs query.IPokemonQueryService
}

// UC: 全検索
func (s PokemonMultiSearchService) SearchPokemons(cmd command.SearchPokemonCommand) (*pokemons.Pokemons, error) {
	return nil, nil
}
