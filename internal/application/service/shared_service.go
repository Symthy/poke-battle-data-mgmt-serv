package service

import (
	"github.com/Symthy/PokeRest/internal/application/command"
	"github.com/Symthy/PokeRest/internal/domain/entity"
	"github.com/Symthy/PokeRest/internal/domain/repository"
)

type PokemonStatsFinder[L entity.IDomains[T, K, uint16], T entity.IDomain[K, uint16], K entity.IValueId[uint16]] struct {
	repo repository.IPokemonStatsRepository[L, T, K]
}

func NewPokemonStatsFinder[L entity.IDomains[T, K, uint16], T entity.IDomain[K, uint16], K entity.IValueId[uint16]](repo repository.IPokemonStatsRepository[L, T, K]) PokemonStatsFinder[L, T, K] {
	return PokemonStatsFinder[L, T, K]{
		repo: repo,
	}
}

func (f PokemonStatsFinder[L, T, K]) FindOfPokemon(pokemonId uint64) (*L, error) {
	// Todo: error when empty
	// Todo: validate id upper limit
	return f.repo.FindOfPokemon(uint16(pokemonId))
}

type EntityAllFinder[L entity.IDomains[T, K, I], T entity.IDomain[K, I], K entity.IValueId[I], I uint16 | uint32 | uint64] struct {
	repo repository.IAllRecordRepository[L, T, K, I]
}

func NewEntityAllFinder[L entity.IDomains[T, K, I], T entity.IDomain[K, I], K entity.IValueId[I], I uint16 | uint32 | uint64](
	repo repository.IAllRecordRepository[L, T, K, I]) EntityAllFinder[L, T, K, I] {
	return EntityAllFinder[L, T, K, I]{
		repo: repo,
	}
}

func (f EntityAllFinder[L, T, K, I]) FindAll(cmd command.PagenationCommand) (*L, error) {
	// Todo: error when empty
	// Todo: error wrap
	return f.repo.FindAll(cmd.Next(), cmd.PageSize())
}
