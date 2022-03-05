package transaction

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

var _ repository.ITrainedPokemonTransactionalRepository = (*TrainedPokemonTransactionalRepository)(nil)

type TrainedPokemonTransactionalRepository struct {
	TransactionalRepositoryWrapper[trainings.TrainedPokemonParam, identifier.TrainedPokemonId]
	repo infrastructure.ISingleRecordFinder[trainings.TrainedPokemonParam, identifier.TrainedPokemonId]
}

func NewTrainedPokemonTransactionalRepository(
	repo repository.ITrainedPokemonRepository,
	dbClient orm.IDbClient,
) *TrainedPokemonTransactionalRepository {
	var innerWriteRepository InnerWriteRepository[trainings.TrainedPokemonParam, identifier.TrainedPokemonId] = repo
	return &TrainedPokemonTransactionalRepository{
		TransactionalRepositoryWrapper: NewTransactionalRepositoryWrapper(innerWriteRepository, dbClient),
	}
}

func (r TrainedPokemonTransactionalRepository) FindById(id uint) (*trainings.TrainedPokemonParam, error) {
	return r.repo.FindById(id)
}
