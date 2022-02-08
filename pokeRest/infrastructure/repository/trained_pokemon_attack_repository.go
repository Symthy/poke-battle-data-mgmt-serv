package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

var _ repository.ITrainedPokemonAttackRepository = (*TrainedPokemonAttackRepository)(nil)

type schemaTpa = schema.TrainedPokemonAttackTarget

var emptyTrainedPokemonAttackSchemaBuilder = func() schemaTpa { return schemaTpa{} }

type TrainedPokemonAttackRepository struct {
	BaseWriteRepository[schemaTpa, pokemons.TrainedPokemonAttackTarget]
	dbClient orm.IDbClient
}

func NewTrainedPokemonAttackRepository(dbClient orm.IDbClient) *TrainedPokemonAttackRepository {
	return &TrainedPokemonAttackRepository{
		BaseWriteRepository: BaseWriteRepository[schemaTpa, pokemons.TrainedPokemonAttackTarget]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyTrainedPokemonAttackSchemaBuilder,
			schemaConverter:    infrastructure.ToSchemaTrainedPokemonAttackTarget,
		},
		dbClient: dbClient,
	}
}
