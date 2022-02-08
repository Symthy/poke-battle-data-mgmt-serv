package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

var _ repository.IPartySeasonResultRepository = (*PartySeasonResultRepository)(nil)

var emptyPartySeasonResultSchemaBuilder = func() schema.PartySeasonResult { return schema.PartySeasonResult{} }

type PartySeasonResultRepository struct {
	BaseWriteRepository[schema.PartySeasonResult, parties.PartySeasonResult]
	dbClient orm.IDbClient
}

func NewPartySeasonResultRepository(dbClient orm.IDbClient) *PartySeasonResultRepository {
	return &PartySeasonResultRepository{
		BaseWriteRepository: BaseWriteRepository[schema.PartySeasonResult, parties.PartySeasonResult]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyPartySeasonResultSchemaBuilder,
			schemaConverter:    infrastructure.ToSchemaPartySeasonResult,
		},
		dbClient: dbClient,
	}
}
