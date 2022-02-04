package mock

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/test/data"
)

// implements IPokemonRepository
type PokemonRepositoryMock struct {
}

func NewPokemonRepositoryMock() *PokemonRepositoryMock {
	return &PokemonRepositoryMock{}
}

func (mock PokemonRepositoryMock) FindById(id uint) (*model.Pokemon, error) {
	dummyPokemon := data.DummyPokemon3()
	if dummyPokemon.ID != id {
		return &model.Pokemon{}, nil
	}
	p := dummyPokemon.ConvertToDomain()
	return &p, nil
}

func (mock PokemonRepositoryMock) FindAll() (model.PokemonList, error) {
	return model.PokemonList{}, nil
}

func (mock PokemonRepositoryMock) Create(pokemon *model.Pokemon) (*model.Pokemon, error) {
	return &model.Pokemon{}, nil
}
