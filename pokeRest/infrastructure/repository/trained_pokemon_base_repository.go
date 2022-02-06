package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.ITrainedPokemonBaseRepository = (*TrainedPokemonBaseRepository)(nil)

type TrainedPokemonBaseRepository struct {
	dbClient orm.IDbClient
}

func (rep TrainedPokemonBaseRepository) Find() {
}

func (rep TrainedPokemonBaseRepository) Save() {
}

func (rep TrainedPokemonBaseRepository) Delete() {
}
