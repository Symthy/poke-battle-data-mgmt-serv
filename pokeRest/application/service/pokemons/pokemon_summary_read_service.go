package pokemons

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type ps = pokemons.Pokemons
type p = pokemons.Pokemon

type PokemonSummaryReadService struct {
	repo repository.IPokemonRepository
	service.EntityAllFinder[ps, p]
}

func NewPokemonSummaryReadService(repo repository.IPokemonRepository) PokemonSummaryReadService {
	serv := PokemonSummaryReadService{repo: repo}
	serv.EntityAllFinder = service.NewEntityAllFinder[ps, p](repo)
	return serv
}

// UC: 単体取得 (need?)
func (s PokemonSummaryReadService) FindPokemon(id uint) (pokemons.Pokemon, error) {
	pokemon, err := s.repo.FindById(id)
	return *pokemon, err
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
