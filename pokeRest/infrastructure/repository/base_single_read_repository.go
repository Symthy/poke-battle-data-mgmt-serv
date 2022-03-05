package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

type BaseSingleReadRepository[TS infrastructure.ISchema[TD, K], TD infrastructure.IDomain[K], K infrastructure.IValueId] struct {
	dbClient           orm.IDbClient
	emptySchemaBuilder func() TS
	toSchemaConverter  func(model TD) TS
	toDomainConverter  func(schema TS) (*TD, error)
}

func (rep BaseSingleReadRepository[TS, TD, K]) FindById(id uint) (*TD, error) {
	db := rep.dbClient.Db()
	schema := rep.emptySchemaBuilder()
	tx := db.First(&schema, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rep.toDomainConverter(schema)
}
