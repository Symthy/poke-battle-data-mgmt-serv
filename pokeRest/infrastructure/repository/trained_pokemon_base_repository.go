package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.ITrainedPokemonBaseRepository = (*TrainedPokemonBaseRepository)(nil)

type TrainedPokemonBaseRepository struct {
	dbClient orm.IDbClient
}

// Todo: args
func (rep TrainedPokemonBaseRepository) Find() (*pokemons.TrainedPokemonBase, error) {
	return nil, nil
}

func (rep TrainedPokemonBaseRepository) Create(model pokemons.TrainedPokemonBase) (*pokemons.TrainedPokemonBase, error) {
	return nil, nil
}

func (rep TrainedPokemonBaseRepository) Update(model pokemons.TrainedPokemonBase) (*pokemons.TrainedPokemonBase, error) {
	return nil, nil
}

func (rep TrainedPokemonBaseRepository) Delete(model pokemons.TrainedPokemonBase) (*pokemons.TrainedPokemonBase, error) {
	return nil, nil
}
