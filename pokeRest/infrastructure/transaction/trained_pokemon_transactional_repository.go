package transaction

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
)

type TrainedPokemonTransactionalRepository struct {
	TransactionRepositoryWrapper[trainings.TrainedPokemonParam]
}

func NewTrainedPokemonTransactionalRepository(
	repo InnerWriteRepository[trainings.TrainedPokemonParam],
	dbClient orm.IDbClient,
) *TrainedPokemonTransactionalRepository {
	return &TrainedPokemonTransactionalRepository{
		TransactionRepositoryWrapper: NewTransactionalRepositoryWrapper(repo, dbClient),
	}
}
