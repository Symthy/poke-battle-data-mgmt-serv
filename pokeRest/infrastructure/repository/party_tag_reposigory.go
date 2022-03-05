package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/conv"
)

var _ repository.IPartyTagRepository = (*PartyTagRepository)(nil)

var (
	emptyPartyTagBuilder  = func() schema.PartyTag { return schema.PartyTag{} }
	emptyPartyTagsBuilder = func() []schema.PartyTag { return []schema.PartyTag{} }
)

type PartyTagRepository struct {
	BaseReadRepository[schema.PartyTag, parties.PartyTag, parties.PartyTags, identifier.PartyTagId]
	BaseWriteRepository[schema.PartyTag, parties.PartyTag, identifier.PartyTagId]
	dbClient orm.IDbClient
}

func NewPartyTagRepository(dbClient orm.IDbClient) *PartyTagRepository {
	return &PartyTagRepository{
		BaseReadRepository: NewBaseReadRepository[schema.PartyTag, parties.PartyTag, parties.PartyTags, identifier.PartyTagId](
			dbClient,
			emptyPartyTagBuilder,
			emptyPartyTagsBuilder,
			parties.NewPartyTags,
			conv.ToSchemaPartyTag,
			conv.ToDomainPartyTag,
		),
		BaseWriteRepository: BaseWriteRepository[schema.PartyTag, parties.PartyTag, identifier.PartyTagId]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyPartyTagBuilder,
			toSchemaConverter:  conv.ToSchemaPartyTag,
			toDomainConverter:  conv.ToDomainPartyTag,
		},
		dbClient: dbClient,
	}
}

// FindById <- BaseReadRepository
// FindAll <- BaseReadRepository

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository
