package service

import (
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type PokemonStatsFinder[L entity.IDomains[T], T entity.IDomain] struct {
	repo repository.IPokemonStatsRepository[L, T]
}

func NewPokemonStatsFinder[L entity.IDomains[T], T entity.IDomain](repo repository.IPokemonStatsRepository[L, T]) PokemonStatsFinder[L, T] {
	return PokemonStatsFinder[L, T]{
		repo: repo,
	}
}

func (f PokemonStatsFinder[L, T]) FindOfPokemon(pokemonId uint) (*L, error) {
	// Todo: error when empty
	// Todo: error wrap
	return f.repo.FindOfPokemon(pokemonId)
}

type EntityAllFinder[L entity.IDomains[T], T entity.IDomain] struct {
	repo repository.IAllRecordRepository[L, T]
}

func NewEntityAllFinder[L entity.IDomains[T], T entity.IDomain](repo repository.IAllRecordRepository[L, T]) EntityAllFinder[L, T] {
	return EntityAllFinder[L, T]{
		repo: repo,
	}
}

func (f EntityAllFinder[L, T]) FindAll(cmd command.PagenationCommand) (*L, error) {
	// Todo: error when empty
	// Todo: error wrap
	return f.repo.FindAll(cmd.Next(), cmd.PageSize())
}
