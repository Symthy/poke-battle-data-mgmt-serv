package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/field"
	"gorm.io/gorm"
)

type BaseReadRepository[TS infrastructure.ISchema[TD], TD infrastructure.IDomain] struct {
	dbClient           orm.IDbClient
	emptySchemaBuilder func() TS
	schemaConverter    func(model TD) TS
}

// Todo: error handling

func (rep BaseReadRepository[TS, TD]) FindById(id uint) (*TD, error) {
	db := rep.dbClient.Db()
	schema := rep.emptySchemaBuilder()
	tx := db.First(&schema, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	s := schema.ConvertToDomain()
	return &s, tx.Error
}

// required: wrap when used
// Todo: return array?
func (rep BaseReadRepository[TS, TD]) FindByField(targetField string, value string, filterFields ...string) (*TD, error) {
	db := rep.dbClient.Db()
	schema := rep.emptySchemaBuilder()
	selectedDbFields := field.ConvertToDbField(filterFields...)

	// Todo: refactor
	var tx *gorm.DB
	if len(selectedDbFields) > 0 {
		tx = db.Select(selectedDbFields).Find(&schema, targetField+" = ?", value)
	} else {
		tx = db.Find(&schema, targetField+" = ?", value)
	}

	if tx.Error != nil {
		return nil, tx.Error
	}
	s := schema.ConvertToDomain()
	return &s, nil
}
