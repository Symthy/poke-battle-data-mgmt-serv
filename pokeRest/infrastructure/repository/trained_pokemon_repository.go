package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.ITrainedPokemonRepository = (*TrainedPokemonRepository)(nil)

type TrainedPokemonRepository struct {
	dbClient orm.IDbClient
}

func (s TrainedPokemonRepository) Create(model pokemons.TrainedPokemon) (*pokemons.TrainedPokemon, error) {
	return nil, nil
}

func (s TrainedPokemonRepository) Update(model pokemons.TrainedPokemon) (*pokemons.TrainedPokemon, error) {
	return nil, nil
}

func (s TrainedPokemonRepository) Delete(model pokemons.TrainedPokemon) (*pokemons.TrainedPokemon, error) {
	return nil, nil
}
