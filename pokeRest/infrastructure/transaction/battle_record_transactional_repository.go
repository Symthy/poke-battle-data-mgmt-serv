package transaction

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

var _ repository.IBattleRecordTransactionalRepository = (*BattleRecordTransactionalRepository)(nil)

type BattleRecordTransactionalRepository struct {
	TransactionalRepositoryWrapper[battles.BattleRecord]
	repo infrastructure.ISingleRecordFinder[battles.BattleRecord, identifier.BattleRecordId]
}

func NewBattleRecordTransactionalRepository(
	repo repository.IBattleRecordRepository,
	dbClient orm.IDbClient,
) *BattleRecordTransactionalRepository {
	var innerWriteRepository InnerWriteRepository[battles.BattleRecord] = repo
	return &BattleRecordTransactionalRepository{
		TransactionalRepositoryWrapper: NewTransactionalRepositoryWrapper(innerWriteRepository, dbClient),
		repo:                           repo,
	}
}

func (r BattleRecordTransactionalRepository) FindById(id uint) (*battles.BattleRecord, error) {
	return r.repo.FindById(id)
}
