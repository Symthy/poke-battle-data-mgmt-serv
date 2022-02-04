package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/common/collections"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
)

type TagRepository struct {
	BaseRepository[schema.Tag, model.Tag]
	dbClient orm.IDbClient
}

func (rep TagRepository) FindAll() ([]model.Tag, error) {
	db := rep.dbClient.Db()
	var schemaTags = []schema.Tag{}

	paginate := rep.dbClient.Paginate(1, 100)
	tx := db.Scopes(paginate).Find(&schemaTags)

	if tx.Error != nil {
		return []model.Tag{}, tx.Error
	}
	tags := collections.ListMap[schema.Tag, model.Tag](schemaTags)
	return tags, nil
}
