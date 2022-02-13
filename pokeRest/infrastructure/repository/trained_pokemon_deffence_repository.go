package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
)

var _ repository.ITrainedPokemonDeffenceRepository = (*TrainedPokemonDeffenceRepository)(nil)

type schemaTpd = schema.TrainedPokemonDeffenceTarget

var emptyTrainedPokemonDeffenceSchemaBuilder = func() schemaTpd { return schemaTpd{} }

type TrainedPokemonDeffenceRepository struct {
	BaseWriteRepository[schemaTpd, trainings.TrainedPokemonDeffenceTarget]
	dbClient orm.IDbClient
}

func NewTrainedPokemonDeffenceRepository(dbClient orm.IDbClient) *TrainedPokemonDeffenceRepository {
	return &TrainedPokemonDeffenceRepository{
		BaseWriteRepository: BaseWriteRepository[schemaTpd, trainings.TrainedPokemonDeffenceTarget]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyTrainedPokemonDeffenceSchemaBuilder,
			schemaConverter:    dto.ToSchemaTrainedPokemonDeffenceTarget,
		},
		dbClient: dbClient,
	}
}

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository
