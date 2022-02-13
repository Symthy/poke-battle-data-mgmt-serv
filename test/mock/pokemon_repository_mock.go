package mock

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/test/data"
)

var _ repository.IPokemonRepository = (*PokemonRepositoryMock)(nil)

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

func (mock PokemonRepositoryMock) FindAll(page int, pageSize int) (*pokemons.Pokemons, error) {
	return &pokemons.Pokemons{}, nil
}

func (mock PokemonRepositoryMock) Create(pokemon pokemons.Pokemon) (*pokemons.Pokemon, error) {
	return &pokemons.Pokemon{}, nil
}

func (mock PokemonRepositoryMock) Update(pokemon pokemons.Pokemon) (*pokemons.Pokemon, error) {
	return &pokemons.Pokemon{}, nil
}

func (mock PokemonRepositoryMock) Delete(id uint) (*pokemons.Pokemon, error) {
	return &pokemons.Pokemon{}, nil
}
