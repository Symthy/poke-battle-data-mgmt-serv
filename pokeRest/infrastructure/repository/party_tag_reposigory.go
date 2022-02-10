package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

var _ repository.IPartyTagRepository = (*PartyTagRepository)(nil)

var emptyPartyTagBuilder = func() schema.PartyTag { return schema.PartyTag{} }

type PartyTagRepository struct {
	BaseReadRepository[schema.PartyTag, parties.PartyTag]
	BaseWriteRepository[schema.PartyTag, parties.PartyTag]
	dbClient orm.IDbClient
}

func NewPartyTagRepository(dbClient orm.IDbClient) *PartyTagRepository {
	return &PartyTagRepository{
		BaseReadRepository: BaseReadRepository[schema.PartyTag, parties.PartyTag]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyPartyTagBuilder,
			schemaConverter:    infrastructure.ToSchemaPartyTag,
		},
		BaseWriteRepository: BaseWriteRepository[schema.PartyTag, parties.PartyTag]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyPartyTagBuilder,
			schemaConverter:    infrastructure.ToSchemaPartyTag,
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
	ts := parties.NewPartyTags(infrastructure.ConvertToDomains[schema.PartyTag, parties.PartyTag](schemaTags))
	return &ts, nil
}
