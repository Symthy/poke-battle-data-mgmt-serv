package repository

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/trainings"
	"github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
	"github.com/Symthy/PokeRest/internal/infrastructure/repository/conv"
)

var _ repository.ITrainedPokemonDeffenceRepository = (*TrainedPokemonDefenceRepository)(nil)

type schemaTpd = schema.TrainedPokemonDefenseTarget

var emptyTrainedPokemonDeffenceSchemaBuilder = func() *schemaTpd { return &schemaTpd{} }

type TrainedPokemonDefenceRepository struct {
	BaseWriteRepository[schemaTpd, trainings.TrainedPokemonDefenseTarget, identifier.TrainedDefenseTargetId, uint64]
	dbClient orm.IDbClient
}

func NewTrainedPokemonDefenceRepository(dbClient orm.IDbClient) *TrainedPokemonDefenceRepository {
	return &TrainedPokemonDefenceRepository{
		BaseWriteRepository: BaseWriteRepository[schemaTpd, trainings.TrainedPokemonDefenseTarget, identifier.TrainedDefenseTargetId, uint64]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyTrainedPokemonDeffenceSchemaBuilder,
			toSchemaConverter:  conv.ToSchemaTrainedPokemonDefenceTarget,
			toDomainConverter:  conv.ToDomainTrainedPokemonDefenceTarget,
		},
		dbClient: dbClient,
	}
}

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository
