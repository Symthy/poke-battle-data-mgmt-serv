package repository

import (
	"fmt"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/field"
	"gorm.io/gorm"
)

type BaseReadRepository[TS infrastructure.ISchema[TD], TD infrastructure.IDomain, TM infrastructure.IDomains[TD]] struct {
	BaseSingleReadRepository[TS, TD]
	dbClient            orm.IDbClient
	emptySchemaBuilder  func() TS
	emptySchemasBuilder func() []TS
	domainsConstructor  func([]TD) TM
	schemaConverter     func(model TD) TS
}

// Todo: error handling

// required: wrap when used
func (rep BaseReadRepository[TS, TD, TM]) FindByField(targetField string, value string, filterFields ...string) (*TM, error) {
	db := rep.dbClient.Db()
	schemas := rep.emptySchemasBuilder()
	selectedDbFields := field.ConvertToDbField(filterFields...)

	// Todo: refactor
	var tx *gorm.DB
	if len(selectedDbFields) > 0 {
		tx = db.Select(selectedDbFields).Find(&schemas, targetField+" = ?", value)
	} else {
		tx = db.Find(&schemas, targetField+" = ?", value)
	}

	if tx.Error != nil {
		return nil, tx.Error
	}
	return rep.resolveReturnValues(schemas)
}

// Todo: args is condition
func (rep BaseReadRepository[TS, TD, TM]) FindAll(next int, pageSize int) (*TM, error) {
	db := rep.dbClient.Db()
	schemas := rep.emptySchemasBuilder()

	paginate := rep.dbClient.Paginate(next, pageSize)
	tx := db.Scopes(paginate).Find(&schemas)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return rep.resolveReturnValues(schemas)
}

func (rep BaseReadRepository[TS, TD, TM]) resolveReturnValues(schemas []TS) (*TM, error) {
	domainArray := dto.ConvertToDomains[TS, TD](schemas)
	domains := dto.BuildDomains(domainArray, rep.domainsConstructor)
	if domains == nil {
		// Todo: error
		return nil, fmt.Errorf("Not found")
	}
	return domains, nil
}
