package transaction

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
)

type TrainedPokemonTransactionalRepository struct {
	TransactionalRepositoryWrapper[trainings.TrainedPokemonParam]
}

func NewTrainedPokemonTransactionalRepository(
	repo InnerWriteRepository[trainings.TrainedPokemonParam],
	dbClient orm.IDbClient,
) *TrainedPokemonTransactionalRepository {
	return &TrainedPokemonTransactionalRepository{
		TransactionalRepositoryWrapper: NewTransactionalRepositoryWrapper(repo, dbClient),
	}
}
