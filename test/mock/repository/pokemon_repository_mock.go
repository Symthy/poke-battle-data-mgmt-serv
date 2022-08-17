package repository

import (
	"github.com/Symthy/PokeRest/internal/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/infrastructure/repository/conv"
	"github.com/Symthy/PokeRest/test/data"
)

var _ repository.IPokemonRepository = (*PokemonRepositoryMock)(nil)

// implements IPokemonRepository
type PokemonRepositoryMock struct {
}

func NewPokemonRepositoryMock() *PokemonRepositoryMock {
	return &PokemonRepositoryMock{}
}

func (mock PokemonRepositoryMock) FindById(id uint16) (*pokemons.Pokemon, error) {
	dummyPokemon := data.DummyPokemon3()
	if dummyPokemon.ID != id {
		return &pokemons.Pokemon{}, nil
	}
	return conv.ToDomainPokemon(dummyPokemon)
}

func (mock PokemonRepositoryMock) FindAll(page uint32, pageSize uint16) (*pokemons.Pokemons, error) {
	return &pokemons.Pokemons{}, nil
}

func (mock PokemonRepositoryMock) Create(pokemon pokemons.Pokemon) (*pokemons.Pokemon, error) {
	return &pokemons.Pokemon{}, nil
}

func (mock PokemonRepositoryMock) Update(pokemon pokemons.Pokemon) (*pokemons.Pokemon, error) {
	return &pokemons.Pokemon{}, nil
}

func (mock PokemonRepositoryMock) Delete(id uint16) (*pokemons.Pokemon, error) {
	return &pokemons.Pokemon{}, nil
}
