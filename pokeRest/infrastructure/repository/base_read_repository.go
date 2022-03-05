package repository

import (
	"fmt"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/conv"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/field"
	"gorm.io/gorm"
)

type BaseReadRepository[TS infrastructure.ISchema[TD, K], TD infrastructure.IDomain[K], TM infrastructure.IDomains[TD, K], K infrastructure.IValueId] struct {
	BaseSingleReadRepository[TS, TD, K]
	dbClient            orm.IDbClient
	emptySchemaBuilder  func() TS
	emptySchemasBuilder func() []TS
	domainsConstructor  func([]TD) TM
	toSchemaConverter   func(TD) TS
}

func NewBaseReadRepository[TS infrastructure.ISchema[TD, K], TD infrastructure.IDomain[K], TM infrastructure.IDomains[TD, K], K infrastructure.IValueId](
	dbClient orm.IDbClient,
	emptySchemaBuilder func() TS,
	emptySchemasBuilder func() []TS,
	domainsConstructor func([]TD) TM,
	toSchemaConverter func(TD) TS,
	toDomainConverter func(TS) (*TD, error),
) BaseReadRepository[TS, TD, TM, K] {
	return BaseReadRepository[TS, TD, TM, K]{
		BaseSingleReadRepository: BaseSingleReadRepository[TS, TD, K]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptySchemaBuilder,
			toSchemaConverter:  toSchemaConverter,
			toDomainConverter:  toDomainConverter,
		},
		dbClient:            dbClient,
		emptySchemaBuilder:  emptySchemaBuilder,
		emptySchemasBuilder: emptySchemasBuilder,
		domainsConstructor:  domainsConstructor,
		toSchemaConverter:   toSchemaConverter,
	}
}

// Todo: error handling

// required: wrap when used
func (rep BaseReadRepository[TS, TD, TM, K]) FindByField(targetField string, value string, filterFields ...string) (*TM, error) {
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
func (rep BaseReadRepository[TS, TD, TM, K]) FindAll(next int, pageSize int) (*TM, error) {
	db := rep.dbClient.Db()
	schemas := rep.emptySchemasBuilder()

	paginate := rep.dbClient.Paginate(next, pageSize)
	tx := db.Scopes(paginate).Find(&schemas)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return rep.resolveReturnValues(schemas)
}

func (rep BaseReadRepository[TS, TD, TM, K]) resolveReturnValues(schemas []TS) (*TM, error) {
	domainArray, err := conv.ConvertToDomains[TS, TD, K](schemas, rep.toDomainConverter)
	if err != nil {
		return nil, err
	}
	domains := dto.BuildDomains[TD, TM, K](domainArray, rep.domainsConstructor)
	if domains == nil {
		// Todo: error
		return nil, fmt.Errorf("Not found")
	}
	return domains, nil
}
