package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
)

var _ repository.IPartyRepository = (*PartyRepository)(nil)

var emptyPartySchemaBuilder = func() schema.Party { return schema.Party{} }

type PartyRepository struct {
	BaseSingleReadRepository[schema.Party, parties.Party, identifier.PartyId]
	BaseWriteRepository[schema.Party, parties.Party, identifier.PartyId]
	dbClient orm.IDbClient
}

func NewPartyRepository(dbClient orm.IDbClient) *PartyRepository {
	return &PartyRepository{
		BaseSingleReadRepository: BaseSingleReadRepository[schema.Party, parties.Party, identifier.PartyId]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyPartySchemaBuilder,
			schemaConverter:    dto.ToSchemaParty,
		},
		BaseWriteRepository: BaseWriteRepository[schema.Party, parties.Party, identifier.PartyId]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyPartySchemaBuilder,
			schemaConverter:    dto.ToSchemaParty,
		},
		dbClient: dbClient,
	}
}

// FindById <- BaseSingleReadRepository

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository
