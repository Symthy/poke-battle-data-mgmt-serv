package transaction

import (
	"fmt"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
	"gorm.io/gorm"
)

var _ IWritableRepository[infrastructure.IDomain] = (*TransactionRepositoryWrapper[infrastructure.IDomain])(nil)

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

type TransactionRepositoryWrapper[TD infrastructure.IDomain] struct {
	InnerWriteRepository[TD]
	dbClient orm.IDbClient
	tx       *gorm.DB
}

func NewTransactionRepositoryWrapper[TD infrastructure.IDomain](
	repo InnerWriteRepository[TD], dbClient orm.IDbClient) TransactionRepositoryWrapper[TD] {
	return TransactionRepositoryWrapper[TD]{
		InnerWriteRepository: repo,
		dbClient:             dbClient,
	}
}

func (trw TransactionRepositoryWrapper[TD]) StartTransaction() error {
	db := trw.dbClient.Db()
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	trw.tx = tx
	return nil
}

func (trw TransactionRepositoryWrapper[TD]) CancelTransaction() error {
	if trw.tx == nil {
		// Todo: error
		return fmt.Errorf("rollback failure")
	}
	err := trw.tx.Rollback().Error
	trw.tx = nil
	return err
}

func (trw TransactionRepositoryWrapper[TD]) FinishTransaction() error {
	if trw.tx == nil {
		// Todo: error
		return fmt.Errorf("commit failure")
	}
	return trw.tx.Commit().Error
}

func (trw TransactionRepositoryWrapper[TD]) PanicPostProcess() {
	if r := recover(); r != nil {
		trw.CancelTransaction()
	}
}

func (trw TransactionRepositoryWrapper[TD]) Create(model TD) (*TD, error) {
	trw.validateTransactioning()
	return trw.CreateRecord(trw.tx, model)
}

func (trw TransactionRepositoryWrapper[TD]) Update(model TD) (*TD, error) {
	trw.validateTransactioning()
	return trw.UpdateRecord(trw.tx, model)
}

func (trw TransactionRepositoryWrapper[TD]) Delete(id uint) (*TD, error) {
	trw.validateTransactioning()
	return trw.DeleteRecord(trw.tx, id)
}

func (trw TransactionRepositoryWrapper[TD]) validateTransactioning() {
	if trw.tx == nil {
		panic("error: non transactioning")
	}
}
