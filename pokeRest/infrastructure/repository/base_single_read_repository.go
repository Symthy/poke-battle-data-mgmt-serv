package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

type BaseSingleReadRepository[TS infrastructure.ISchema[TD, K, I], TD infrastructure.IDomain[K, I], K infrastructure.IValueId[I], I uint16 | uint32 | uint64] struct {
	dbClient           orm.IDbClient
	emptySchemaBuilder func() TS
	toSchemaConverter  func(model TD) TS
	toDomainConverter  func(schema TS) (*TD, error)
}

func (rep BaseSingleReadRepository[TS, TD, K, I]) FindById(id I) (*TD, error) {
	db := rep.dbClient.Db()
	schema := rep.emptySchemaBuilder()
	tx := db.First(&schema, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rep.toDomainConverter(schema)
}
