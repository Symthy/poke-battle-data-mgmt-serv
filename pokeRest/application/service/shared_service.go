package service

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type StatsOfPokemonFinder[L model.IDomains[T], T model.IDomain] struct {
	repo repository.IValueOfPokemonRepository[L, T]
}

func NewStatsOfPokemonFinder[L model.IDomains[T], T model.IDomain](repo repository.IValueOfPokemonRepository[L, T]) StatsOfPokemonFinder[L, T] {
	return StatsOfPokemonFinder[L, T]{
		repo: repo,
	}
}

func (s StatsOfPokemonFinder[L, T]) FindOfPokemon(pokemonId uint) (*L, error) {
	return s.repo.FindOfPokemon(pokemonId)
}
