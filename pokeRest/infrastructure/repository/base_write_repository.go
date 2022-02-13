package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

type BaseWriteRepository[TS infrastructure.ISchema[TD], TD infrastructure.IDomain] struct {
	dbClient           orm.IDbClient
	emptySchemaBuilder func() TS
	schemaConverter    func(model TD) TS
}

// Todo: error handling

func (rep BaseWriteRepository[TS, TD]) Create(model TD) (*TD, error) {
	schema := rep.schemaConverter(model)
	db := rep.dbClient.Db()
	tx := db.Create(&schema)
	if tx.Error != nil {
		return nil, tx.Error
	}
	s := schema.ConvertToDomain()
	return &s, nil
}

func (rep BaseWriteRepository[TS, TD]) Update(model TD) (*TD, error) {
	db := rep.dbClient.Db()
	target := rep.emptySchemaBuilder()
	if tx := db.First(&target, model.Id()); tx.Error != nil {
		return nil, tx.Error
	}
	schema := rep.schemaConverter(model)
	tx := db.Model(&target).Updates(&schema)
	if tx.Error != nil {
		return nil, tx.Error
	}
	s := schema.ConvertToDomain()
	return &s, nil
}

func (rep BaseWriteRepository[TS, TD]) Delete(id uint) (*TD, error) {
	db := rep.dbClient.Db()
	if tx := db.First(rep.emptySchemaBuilder(), id); tx.Error != nil {
		return nil, tx.Error
	}
	schema := rep.emptySchemaBuilder()
	tx := db.Delete(&schema, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	s := schema.ConvertToDomain()
	return &s, nil
}
