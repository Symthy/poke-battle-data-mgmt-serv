package repository

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/parties"
	"github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
	"github.com/Symthy/PokeRest/internal/infrastructure/repository/conv"
)

var _ repository.IPartyRepository = (*PartyRepository)(nil)

var emptyPartySchemaBuilder = func() *schema.Party { return &schema.Party{} }

type PartyRepository struct {
	BaseSingleReadRepository[schema.Party, parties.Party, identifier.PartyId, uint64]
	BaseWriteRepository[schema.Party, parties.Party, identifier.PartyId, uint64]
	dbClient orm.IDbClient
}

func NewPartyRepository(dbClient orm.IDbClient) *PartyRepository {
	return &PartyRepository{
		BaseSingleReadRepository: BaseSingleReadRepository[schema.Party, parties.Party, identifier.PartyId, uint64]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyPartySchemaBuilder,
			toSchemaConverter:  conv.ToSchemaParty,
			toDomainConverter:  conv.ToDomainParty,
		},
		BaseWriteRepository: BaseWriteRepository[schema.Party, parties.Party, identifier.PartyId, uint64]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyPartySchemaBuilder,
			toSchemaConverter:  conv.ToSchemaParty,
			toDomainConverter:  conv.ToDomainParty,
		},
		dbClient: dbClient,
	}
}

// FindById <- BaseSingleReadRepository

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository
