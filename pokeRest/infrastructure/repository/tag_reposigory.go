package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/tags"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

var _ repository.ITagRepository = (*TagRepository)(nil)

var emptyTagBuilder = func() schema.Tag { return schema.Tag{} }

type TagRepository struct {
	BaseReadRepository[schema.Tag, tags.Tag]
	BaseWriteRepository[schema.Tag, tags.Tag]
	dbClient orm.IDbClient
}

func NewTagRepository(dbClient orm.IDbClient) *TagRepository {
	return &TagRepository{
		BaseReadRepository: BaseReadRepository[schema.Tag, tags.Tag]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyTagBuilder,
			schemaConverter:    infrastructure.ToSchemaTag,
		},
		BaseWriteRepository: BaseWriteRepository[schema.Tag, tags.Tag]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyTagBuilder,
			schemaConverter:    infrastructure.ToSchemaTag,
		},
		dbClient: dbClient,
	}
}

func (rep TagRepository) FindAll(page int, pageSize int) (*tags.Tags, error) {
	var schemaTags = []schema.Tag{}

	paginate := rep.dbClient.Paginate(page, pageSize)
	tx := rep.dbClient.Db().Scopes(paginate).Find(&schemaTags)

	if tx.Error != nil {
		return nil, tx.Error
	}
	ts := tags.NewTags(infrastructure.ConvertSchemasToDomains[schema.Tag, tags.Tag](schemaTags))
	return &ts, nil
}
