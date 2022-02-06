package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.ITrainedPokemonRepository = (*TrainedPokemonRepository)(nil)

type TrainedPokemonRepository struct {
	dbClient orm.IDbClient
}

func (s TrainedPokemonRepository) Save() {
}

func (s TrainedPokemonRepository) Update() {
}

func (s TrainedPokemonRepository) Delete() {
}
