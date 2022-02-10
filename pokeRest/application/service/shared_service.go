package service

import (
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type PokemonStatsFinder[L model.IDomains[T], T model.IDomain] struct {
	repo repository.IPokemonStatsRepository[L, T]
}

func NewPokemonStatsFinder[L model.IDomains[T], T model.IDomain](repo repository.IPokemonStatsRepository[L, T]) PokemonStatsFinder[L, T] {
	return PokemonStatsFinder[L, T]{
		repo: repo,
	}
}

func (f PokemonStatsFinder[L, T]) FindOfPokemon(pokemonId uint) (*L, error) {
	// Todo: error when empty
	// Todo: error wrap
	return f.repo.FindOfPokemon(pokemonId)
}

type EntityAllFinder[L model.IDomains[T], T model.IDomain] struct {
	repo repository.IEntityAllRepository[L, T]
}

func NewEntityAllFinder[L model.IDomains[T], T model.IDomain](repo repository.IEntityAllRepository[L, T]) EntityAllFinder[L, T] {
	return EntityAllFinder[L, T]{
		repo: repo,
	}
}

func (f EntityAllFinder[L, T]) FindAll(cmd command.GetAllEntityCommand) (*L, error) {
	// Todo: error when empty
	// Todo: error wrap
	return f.repo.FindAll(cmd.Page(), cmd.PageSize())
}
