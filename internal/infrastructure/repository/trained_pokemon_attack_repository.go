package repository

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/trainings"
	"github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
	"github.com/Symthy/PokeRest/internal/infrastructure/repository/conv"
)

var _ repository.ITrainedPokemonAttackRepository = (*TrainedPokemonAttackRepository)(nil)

type schemaTpa = schema.TrainedPokemonAttackTarget

var emptyTrainedPokemonAttackSchemaBuilder = func() *schemaTpa { return &schemaTpa{} }

type TrainedPokemonAttackRepository struct {
	BaseWriteRepository[schemaTpa, trainings.TrainedPokemonAttackTarget, identifier.TrainedAttackTargetId, uint64]
	dbClient orm.IDbClient
}

func NewTrainedPokemonAttackRepository(dbClient orm.IDbClient) *TrainedPokemonAttackRepository {
	return &TrainedPokemonAttackRepository{
		BaseWriteRepository: BaseWriteRepository[schemaTpa, trainings.TrainedPokemonAttackTarget, identifier.TrainedAttackTargetId, uint64]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyTrainedPokemonAttackSchemaBuilder,
			toSchemaConverter:  conv.ToSchemaTrainedPokemonAttackTarget,
			toDomainConverter:  conv.ToDomainTrainedPokemonAttackTarget,
		},
		dbClient: dbClient,
	}
}

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository
