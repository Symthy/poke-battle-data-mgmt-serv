package transaction

import (
	"fmt"

	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/infrastructure"
	"gorm.io/gorm"
)

var _ IWritableRepository[
	infrastructure.IDomain[infrastructure.IValueId[uint64], uint64], infrastructure.IValueId[uint64],
] = (*TransactionalRepositoryWrapper[
	infrastructure.IDomain[infrastructure.IValueId[uint64], uint64],
	infrastructure.IValueId[uint64],
])(nil)

type InnerWriteRepository[TD infrastructure.IDomain[K, uint64], K infrastructure.IValueId[uint64]] interface {
	CreateRecord(*gorm.DB, TD) (*TD, error)
	UpdateRecord(*gorm.DB, TD) (*TD, error)
	DeleteRecord(*gorm.DB, uint) (*TD, error)
}

type IWritableRepository[TD infrastructure.IDomain[K, uint64], K infrastructure.IValueId[uint64]] interface {
	Create(TD) (*TD, error)
	Update(TD) (*TD, error)
	Delete(uint) (*TD, error)
}

type TransactionalRepositoryWrapper[TD infrastructure.IDomain[K, uint64], K infrastructure.IValueId[uint64]] struct {
	InnerWriteRepository[TD, K]
	dbClient orm.IDbClient
	tx       *gorm.DB
}

func NewTransactionalRepositoryWrapper[TD infrastructure.IDomain[K, uint64], K infrastructure.IValueId[uint64]](
	repo InnerWriteRepository[TD, K], dbClient orm.IDbClient) TransactionalRepositoryWrapper[TD, K] {
	return TransactionalRepositoryWrapper[TD, K]{
		InnerWriteRepository: repo,
		dbClient:             dbClient,
	}
}

func (trw *TransactionalRepositoryWrapper[TD, K]) StartTransaction() error {
	db := trw.dbClient.Db()
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	trw.tx = tx
	return nil
}

func (trw *TransactionalRepositoryWrapper[TD, K]) CancelTransaction() error {
	if trw.tx == nil {
		// Todo: error
		return fmt.Errorf("could't rollback")
	}
	err := trw.tx.Rollback().Error
	trw.tx = nil
	return err
}

func (trw *TransactionalRepositoryWrapper[TD, K]) FinishTransaction() error {
	if trw.tx == nil {
		// Todo: error
		return fmt.Errorf("could't commit")
	}
	return trw.tx.Commit().Error
}

func (trw *TransactionalRepositoryWrapper[TD, K]) PanicPostProcess() error {
	if r := recover(); r != nil {
		return trw.CancelTransaction()
	}
	return nil
}

func (trw *TransactionalRepositoryWrapper[TD, K]) Create(model TD) (*TD, error) {
	trw.validateTransactioning()
	return trw.CreateRecord(trw.tx, model)
}

func (trw *TransactionalRepositoryWrapper[TD, K]) Update(model TD) (*TD, error) {
	trw.validateTransactioning()
	return trw.UpdateRecord(trw.tx, model)
}

func (trw *TransactionalRepositoryWrapper[TD, K]) Delete(id uint) (*TD, error) {
	trw.validateTransactioning()
	return trw.DeleteRecord(trw.tx, id)
}

func (trw *TransactionalRepositoryWrapper[TD, K]) validateTransactioning() {
	if trw.tx == nil {
		panic("error: non transactioning")
	}
}
