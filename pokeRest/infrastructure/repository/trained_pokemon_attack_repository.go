package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
)

var _ repository.ITrainedPokemonAttackRepository = (*TrainedPokemonAttackRepository)(nil)

type schemaTpa = schema.TrainedPokemonAttackTarget

var emptyTrainedPokemonAttackSchemaBuilder = func() schemaTpa { return schemaTpa{} }

type TrainedPokemonAttackRepository struct {
	BaseWriteRepository[schemaTpa, trainings.TrainedPokemonAttackTarget, identifier.TrainedAttackTargetId]
	dbClient orm.IDbClient
}

func NewTrainedPokemonAttackRepository(dbClient orm.IDbClient) *TrainedPokemonAttackRepository {
	return &TrainedPokemonAttackRepository{
		BaseWriteRepository: BaseWriteRepository[schemaTpa, trainings.TrainedPokemonAttackTarget, identifier.TrainedAttackTargetId]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyTrainedPokemonAttackSchemaBuilder,
			schemaConverter:    dto.ToSchemaTrainedPokemonAttackTarget,
		},
		dbClient: dbClient,
	}
}

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository
