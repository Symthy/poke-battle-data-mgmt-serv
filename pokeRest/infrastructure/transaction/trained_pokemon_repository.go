package transaction

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
)

type TrainedPokemonRepositoryWrapper struct {
	TransactionRepositoryWrapper[trainings.TrainedPokemonParam]
}

func NewTrainedPokemonRepositoryWrapper(
	repo InnerWriteRepository[trainings.TrainedPokemonParam],
	dbClient orm.IDbClient,
) TrainedPokemonRepositoryWrapper {
	return TrainedPokemonRepositoryWrapper{
		TransactionRepositoryWrapper: NewTransactionRepositoryWrapper(repo, dbClient),
	}
}
