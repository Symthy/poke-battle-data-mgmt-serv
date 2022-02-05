package mock

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
	"github.com/Symthy/PokeRest/test/data"
)

// implements IPokemonRepository
type PokemonRepositoryMock struct {
}

func NewPokemonRepositoryMock() *PokemonRepositoryMock {
	return &PokemonRepositoryMock{}
}

func (mock PokemonRepositoryMock) FindById(id uint) (*pokemons.Pokemon, error) {
	dummyPokemon := data.DummyPokemon3()
	if dummyPokemon.ID != id {
		return &pokemons.Pokemon{}, nil
	}
	p := dummyPokemon.ConvertToDomain()
	return &p, nil
}

func (mock PokemonRepositoryMock) FindAll() (pokemons.PokemonList, error) {
	return pokemons.PokemonList{}, nil
}

func (mock PokemonRepositoryMock) Create(pokemon *pokemons.Pokemon) (*pokemons.Pokemon, error) {
	return &pokemons.Pokemon{}, nil
}
