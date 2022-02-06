package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/common/collections"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/tags"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.ITagRepository = (*TagRepository)(nil)

type TagRepository struct {
	BaseRepository[schema.Tag, tags.Tag]
	dbClient orm.IDbClient
}

func NewTagRepository(dbClient orm.IDbClient) TagRepository {
	return TagRepository{
		BaseRepository: BaseRepository[schema.Tag, tags.Tag]{
			dbClient:           dbClient,
			emptySchemaBuilder: func() schema.Tag { return schema.Tag{} },
			schemaConverter:    orm.ToSchemaTag,
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
	ts := tags.NewTags(collections.ListMap[schema.Tag, tags.Tag](schemaTags))
	return &ts, nil
}
