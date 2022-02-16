package transaction

import (
	"fmt"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
	"gorm.io/gorm"
)

var _ IWritableRepository[infrastructure.IDomain] = (*TransactionalRepositoryWrapper[infrastructure.IDomain])(nil)

type InnerWriteRepository[TD infrastructure.IDomain] interface {
	CreateRecord(*gorm.DB, TD) (*TD, error)
	UpdateRecord(*gorm.DB, TD) (*TD, error)
	DeleteRecord(*gorm.DB, uint) (*TD, error)
}

type IWritableRepository[TD infrastructure.IDomain] interface {
	Create(TD) (*TD, error)
	Update(TD) (*TD, error)
	Delete(uint) (*TD, error)
}

type TransactionalRepositoryWrapper[TD infrastructure.IDomain] struct {
	InnerWriteRepository[TD]
	dbClient orm.IDbClient
	tx       *gorm.DB
}

func NewTransactionalRepositoryWrapper[TD infrastructure.IDomain](
	repo InnerWriteRepository[TD], dbClient orm.IDbClient) TransactionalRepositoryWrapper[TD] {
	return TransactionalRepositoryWrapper[TD]{
		InnerWriteRepository: repo,
		dbClient:             dbClient,
	}
}

func (trw TransactionalRepositoryWrapper[TD]) StartTransaction() error {
	db := trw.dbClient.Db()
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	trw.tx = tx
	return nil
}

func (trw TransactionalRepositoryWrapper[TD]) CancelTransaction() error {
	if trw.tx == nil {
		// Todo: error
		return fmt.Errorf("rollback failure")
	}
	err := trw.tx.Rollback().Error
	trw.tx = nil
	return err
}

func (trw TransactionalRepositoryWrapper[TD]) FinishTransaction() error {
	if trw.tx == nil {
		// Todo: error
		return fmt.Errorf("commit failure")
	}
	return trw.tx.Commit().Error
}

func (trw TransactionalRepositoryWrapper[TD]) PanicPostProcess() {
	if r := recover(); r != nil {
		trw.CancelTransaction()
	}
}

func (trw TransactionalRepositoryWrapper[TD]) Create(model TD) (*TD, error) {
	trw.validateTransactioning()
	return trw.CreateRecord(trw.tx, model)
}

func (trw TransactionalRepositoryWrapper[TD]) Update(model TD) (*TD, error) {
	trw.validateTransactioning()
	return trw.UpdateRecord(trw.tx, model)
}

func (trw TransactionalRepositoryWrapper[TD]) Delete(id uint) (*TD, error) {
	trw.validateTransactioning()
	return trw.DeleteRecord(trw.tx, id)
}

func (trw TransactionalRepositoryWrapper[TD]) validateTransactioning() {
	if trw.tx == nil {
		panic("error: non transactioning")
	}
}
