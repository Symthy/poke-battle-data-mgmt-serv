package pokemons

import (
	"github.com/Symthy/PokeRest/internal/application/service"
	"github.com/Symthy/PokeRest/internal/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type ps = pokemons.Pokemons
type p = pokemons.Pokemon
type pi = identifier.PokemonId

type PokemonSummaryReadService struct {
	repo repository.IPokemonRepository
	service.EntityAllFinder[ps, p, pi, uint16]
}

func NewPokemonSummaryReadService(repo repository.IPokemonRepository) PokemonSummaryReadService {
	serv := PokemonSummaryReadService{repo: repo}
	serv.EntityAllFinder = service.NewEntityAllFinder[ps, p, pi, uint16](repo)
	return serv
}

// UC: 単体取得 (need?)
func (s PokemonSummaryReadService) FindPokemon(id uint64) (*pokemons.Pokemon, error) {
	// Todo: validate id upper limit
	pokemon, err := s.repo.FindById(uint16(id))
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
