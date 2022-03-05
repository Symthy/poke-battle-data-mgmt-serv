package service

import (
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type PokemonStatsFinder[L entity.IDomains[T, K], T entity.IDomain[K], K entity.IValueId] struct {
	repo repository.IPokemonStatsRepository[L, T, K]
}

func NewPokemonStatsFinder[L entity.IDomains[T, K], T entity.IDomain[K], K entity.IValueId](repo repository.IPokemonStatsRepository[L, T, K]) PokemonStatsFinder[L, T, K] {
	return PokemonStatsFinder[L, T, K]{
		repo: repo,
	}
}

func (f PokemonStatsFinder[L, T, K]) FindOfPokemon(pokemonId uint) (*L, error) {
	// Todo: error when empty
	// Todo: error wrap
	return f.repo.FindOfPokemon(pokemonId)
}

type EntityAllFinder[L entity.IDomains[T, K], T entity.IDomain[K], K entity.IValueId] struct {
	repo repository.IAllRecordRepository[L, T, K]
}

func NewEntityAllFinder[L entity.IDomains[T, K], T entity.IDomain[K], K entity.IValueId](repo repository.IAllRecordRepository[L, T, K]) EntityAllFinder[L, T, K] {
	return EntityAllFinder[L, T, K]{
		repo: repo,
	}
}

func (f EntityAllFinder[L, T, K]) FindAll(cmd command.PagenationCommand) (*L, error) {
	// Todo: error when empty
	// Todo: error wrap
	return f.repo.FindAll(cmd.Next(), cmd.PageSize())
}
