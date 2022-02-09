package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

var _ repository.ITrainedPokemonMoveSetRepository = (*TrainedPokemonMoveSetRepository)(nil)

type schemaTpb = schema.TrainedPokemonMoveSet

var emptyTrainedPokemonBaseSchemaBuilder = func() schemaTpb { return schemaTpb{} }

type TrainedPokemonMoveSetRepository struct {
	BaseWriteRepository[schemaTpb, pokemons.TrainedPokemonMoveSet]
	dbClient orm.IDbClient
}

func NewTrainedPokemonMoveSetRepository(dbClient orm.IDbClient) *TrainedPokemonMoveSetRepository {
	return &TrainedPokemonMoveSetRepository{
		BaseWriteRepository: BaseWriteRepository[schemaTpb, pokemons.TrainedPokemonMoveSet]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyTrainedPokemonBaseSchemaBuilder,
			schemaConverter:    infrastructure.ToSchemaTrainedPokemonMoveSet,
		},
		dbClient: dbClient,
	}
}

// Todo: args
func (rep TrainedPokemonMoveSetRepository) Find() (*pokemons.TrainedPokemonMoveSet, error) {
	return nil, nil
}
