package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/conv"
)

var _ repository.ITrainedPokemonDeffenceRepository = (*TrainedPokemonDefenceRepository)(nil)

type schemaTpd = schema.TrainedPokemonDefenceTarget

var emptyTrainedPokemonDeffenceSchemaBuilder = func() schemaTpd { return schemaTpd{} }

type TrainedPokemonDefenceRepository struct {
	BaseWriteRepository[schemaTpd, trainings.TrainedPokemonDefenceTarget, identifier.TrainedDefenseTargetId]
	dbClient orm.IDbClient
}

func NewTrainedPokemonDefenceRepository(dbClient orm.IDbClient) *TrainedPokemonDefenceRepository {
	return &TrainedPokemonDefenceRepository{
		BaseWriteRepository: BaseWriteRepository[schemaTpd, trainings.TrainedPokemonDefenceTarget, identifier.TrainedDefenseTargetId]{
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
