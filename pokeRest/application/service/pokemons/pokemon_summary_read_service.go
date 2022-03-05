package pokemons

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type ps = pokemons.Pokemons
type p = pokemons.Pokemon
type pi = identifier.PokemonId

type PokemonSummaryReadService struct {
	repo repository.IPokemonRepository
	service.EntityAllFinder[ps, p, pi]
}

func NewPokemonSummaryReadService(repo repository.IPokemonRepository) PokemonSummaryReadService {
	serv := PokemonSummaryReadService{repo: repo}
	serv.EntityAllFinder = service.NewEntityAllFinder[ps, p, pi](repo)
	return serv
}

// UC: 単体取得 (need?)
func (s PokemonSummaryReadService) FindPokemon(id uint) (*pokemons.Pokemon, error) {
	pokemon, err := s.repo.FindById(id)
	return pokemon, err
}

// UC: 一覧取得 EntityAllFinder

// UC:技による検索
func (s PokemonSummaryReadService) FindByMove(moveId uint) (*pokemons.Pokemons, error) {
	return nil, nil
}

// UC:特性による検索
func (s PokemonSummaryReadService) FindByAbility(abilityId uint) (*pokemons.Pokemons, error) {
	return nil, nil
}
