package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

var _ repository.ITrainedPokemonBaseRepository = (*TrainedPokemonBaseRepository)(nil)

type schemaTpb = schema.TrainedPokemonBase

var emptyTrainedPokemonBaseSchemaBuilder = func() schemaTpb { return schemaTpb{} }

type TrainedPokemonBaseRepository struct {
	BaseWriteRepository[schemaTpb, pokemons.TrainedPokemonBase]
	dbClient orm.IDbClient
}

func NewTrainedPokemonBaseRepository(dbClient orm.IDbClient) *TrainedPokemonBaseRepository {
	return &TrainedPokemonBaseRepository{
		BaseWriteRepository: BaseWriteRepository[schemaTpb, pokemons.TrainedPokemonBase]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyTrainedPokemonBaseSchemaBuilder,
			schemaConverter:    infrastructure.ToSchemaTrainedPokemonBase,
		},
		dbClient: dbClient,
	}
}

// Todo: args
func (rep TrainedPokemonBaseRepository) Find() (*pokemons.TrainedPokemonBase, error) {
	return nil, nil
}
