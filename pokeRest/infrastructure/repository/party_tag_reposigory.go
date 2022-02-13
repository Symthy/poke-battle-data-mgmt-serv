package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
)

var _ repository.IPartyTagRepository = (*PartyTagRepository)(nil)

var (
	emptyPartyTagBuilder  = func() schema.PartyTag { return schema.PartyTag{} }
	emptyPartyTagsBuilder = func() []schema.PartyTag { return []schema.PartyTag{} }
)

type PartyTagRepository struct {
	BaseReadRepository[schema.PartyTag, parties.PartyTag, parties.PartyTags]
	BaseWriteRepository[schema.PartyTag, parties.PartyTag]
	dbClient orm.IDbClient
}

func NewPartyTagRepository(dbClient orm.IDbClient) *PartyTagRepository {
	return &PartyTagRepository{
		BaseReadRepository: BaseReadRepository[schema.PartyTag, parties.PartyTag, parties.PartyTags]{
			dbClient:            dbClient,
			emptySchemaBuilder:  emptyPartyTagBuilder,
			emptySchemasBuilder: emptyPartyTagsBuilder,
			domainsConstructor:  parties.NewPartyTags,
			schemaConverter:     dto.ToSchemaPartyTag,
		},
		BaseWriteRepository: BaseWriteRepository[schema.PartyTag, parties.PartyTag]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyPartyTagBuilder,
			schemaConverter:    dto.ToSchemaPartyTag,
		},
		dbClient: dbClient,
	}
}

func (rep PartyTagRepository) FindAll(page int, pageSize int) (*parties.PartyTags, error) {
	var schemaTags = []schema.PartyTag{}

	paginate := rep.dbClient.Paginate(page, pageSize)
	tx := rep.dbClient.Db().Scopes(paginate).Find(&schemaTags)

	if tx.Error != nil {
		return nil, tx.Error
	}
	ts := parties.NewPartyTags(dto.ConvertToDomains[schema.PartyTag, parties.PartyTag](schemaTags))
	return &ts, nil
}
