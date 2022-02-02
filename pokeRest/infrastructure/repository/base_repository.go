package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/field"
	"gorm.io/gorm"
)

type BaseRepository[TS infrastructure.ISchema[TD], TD infrastructure.IDomain] struct {
	dbClient           orm.IDbClient
	emptySchemaBuilder func() TS
	schemaConverter    func(model TD) TS
}

func (rep BaseRepository[TS, TD]) FindById(id uint) (*TD, error) {
	db := rep.dbClient.Db()
	var schema = rep.emptySchemaBuilder()
	tx := db.Limit(1).Find(&schema, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return schema.ConvertToDomain(), tx.Error
}

func (rep BaseRepository[TS, TD]) FindByField(targetField string, value string, filterFields ...string) (*TD, error) {
	db := rep.dbClient.Db()
	var schema = rep.emptySchemaBuilder()
	selectedDbFields := field.ConvertToDbField(filterFields...)

	var tx *gorm.DB
	if len(selectedDbFields) > 0 {
		tx = db.Select(selectedDbFields).Find(&schema, targetField+" = ?", value)
	} else {
		tx = db.Find(&schema, targetField+" = ?", value)
	}

	if tx.Error != nil {
		return nil, tx.Error
	}
	return schema.ConvertToDomain(), nil
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
