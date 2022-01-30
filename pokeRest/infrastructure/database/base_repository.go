package database

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

// TODO: common repository by generics
type BaseRepository[TS infrastructure.ISchema[TD], TD infrastructure.IDomain] struct {
	dbClient           orm.IDbClient
	emptySchemaBuilder func() TS
	schemaConverter    func(model TD) TS
}

func (rep BaseRepository[TS, TD]) FindById(id uint) (*TD, error) {
	db := rep.dbClient.Db()
	var schema = rep.emptySchemaBuilder()
	tx := db.First(&schema, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return schema.ConvertToDomain(), tx.Error
}

// Todo
func FindAll() {
	return
}

func (rep BaseRepository[TS, TD]) Create(model TD) (*TD, error) {
	schema := rep.schemaConverter(model)
	db := rep.dbClient.Db()
	tx := db.Create(&schema)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return schema.ConvertToDomain(), nil
}

func (rep BaseRepository[TS, TD]) Update(model TD) (*TD, error) {
	db := rep.dbClient.Db()
	target, err := rep.FindById(model.Id())
	if err != nil {
		return nil, err
	}
	schema := rep.schemaConverter(model)
	tx := db.Model(&target).Updates(&schema)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return schema.ConvertToDomain(), nil
}

func (rep BaseRepository[TS, TD]) Delete(id uint) (*TD, error) {
	db := rep.dbClient.Db()
	schema := rep.emptySchemaBuilder()
	tx := db.Delete(&schema, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return schema.ConvertToDomain(), nil
}
