package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

var _ repository.IPartyRepository = (*PartyRepository)(nil)

var emptyPartySchemaBuilder = func() schema.Party { return schema.Party{} }

type PartyRepository struct {
	BaseWriteRepository[schema.Party, parties.Party]
	dbClient orm.IDbClient
}

func NewPartyRepository(dbClient orm.IDbClient) *PartyRepository {
	return &PartyRepository{
		BaseWriteRepository: BaseWriteRepository[schema.Party, parties.Party]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyPartySchemaBuilder,
			schemaConverter:    infrastructure.ToSchemaParty,
		},
		dbClient: dbClient,
	}
}
