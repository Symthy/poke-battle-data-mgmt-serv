package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
	"gorm.io/gorm"
)

type BaseWriteRepository[TS infrastructure.ISchema[TD, K], TD infrastructure.IDomain[K], K infrastructure.IValueId] struct {
	dbClient           orm.IDbClient
	emptySchemaBuilder func() TS
	toSchemaConverter  func(model TD) TS
	toDomainConverter  func(schema TS) (*TD, error)
}

// Todo: error handling

func (rep BaseWriteRepository[TS, TD, K]) Create(model TD) (*TD, error) {
	return rep.CreateRecord(rep.dbClient.Db(), model)
}

func (rep BaseWriteRepository[TS, TD, K]) CreateRecord(db *gorm.DB, model TD) (*TD, error) {
	schema := rep.toSchemaConverter(model)
	tx := db.Create(&schema)
	if tx.Error != nil {
		return nil, tx.Error
	}
	domain, err := rep.toDomainConverter(schema)
	return domain, err
}

func (rep BaseWriteRepository[TS, TD, K]) Update(model TD) (*TD, error) {
	return rep.UpdateRecord(rep.dbClient.Db(), model)
}

func (rep BaseWriteRepository[TS, TD, K]) UpdateRecord(db *gorm.DB, model TD) (*TD, error) {
	target := rep.emptySchemaBuilder()
	if tx := db.First(&target, model.Id().Value()); tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, tx.Error
	}
	schema := rep.toSchemaConverter(model)
	tx := db.Model(&target).Updates(&schema)
	if tx.Error != nil {
		return nil, tx.Error
	}
	domain, err := rep.toDomainConverter(schema)
	return domain, err
}

func (rep BaseWriteRepository[TS, TD, K]) Delete(id uint) (*TD, error) {
	return rep.DeleteRecord(rep.dbClient.Db(), id)
}

func (rep BaseWriteRepository[TS, TD, K]) DeleteRecord(db *gorm.DB, id uint) (*TD, error) {
	if tx := db.First(rep.emptySchemaBuilder(), id); tx.Error != nil {
		return nil, tx.Error
	}
	schema := rep.emptySchemaBuilder()
	tx := db.Delete(&schema, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	domain, err := rep.toDomainConverter(schema)
	return domain, err
}
