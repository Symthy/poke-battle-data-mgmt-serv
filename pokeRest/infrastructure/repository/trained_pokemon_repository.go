package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
)

var _ repository.ITrainedPokemonRepository = (*TrainedPokemonRepository)(nil)

var emptyTrainedPokemonSchemaBuilder = func() schema.TrainedPokemon { return schema.TrainedPokemon{} }

type TrainedPokemonRepository struct {
	BaseWriteRepository[schema.TrainedPokemon, trainings.TrainedPokemonParam, identifier.TrainedPokemonId]
	dbClient orm.IDbClient
}

func NewTrainedPokemonRepository(dbClient orm.IDbClient) *TrainedPokemonRepository {
	return &TrainedPokemonRepository{
		BaseWriteRepository: BaseWriteRepository[schema.TrainedPokemon, trainings.TrainedPokemonParam, identifier.TrainedPokemonId]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyTrainedPokemonSchemaBuilder,
			schemaConverter:    dto.ToSchemaTrainedPokemon,
		},
		dbClient: dbClient,
	}
}

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository
