package battles

import (
	"github.com/Symthy/PokeRest/internal/application/service/battles/command"
	"github.com/Symthy/PokeRest/internal/application/service/battles/spec"
	"github.com/Symthy/PokeRest/internal/domain/entity/battles"
	"github.com/Symthy/PokeRest/internal/domain/repository"
)

type battleRecordtransactionalRepositoryBuilder = func(repository.IBattleRecordRepository) repository.IBattleRecordTransactionalRepository

type BattleRecordWriteService struct {
	battleRecordRepo repository.IBattleRecordRepository
	spec             spec.BattleRecordSpecification
}

func NewBattleRecordWriteService(
	battleRecordRepo repository.IBattleRecordRepository,
	partyRepo repository.IPartyRepository,
) BattleRecordWriteService {
	return BattleRecordWriteService{
		battleRecordRepo: battleRecordRepo,
		spec:             spec.NewBattleRecordSpecification(battleRecordRepo, partyRepo),
	}
}

// UC: 戦績登録 (パーティ戦績も更新)
func (s BattleRecordWriteService) AddBattleRecord(cmd command.AddBattleRecordCommand) (*battles.BattleRecord, error) {
	battleRecord, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	if _, err := s.spec.IsSatisfyForUpdate(battleRecord); err != nil {
		return nil, err
	}
	createdBattleRecord, err := s.battleRecordRepo.Create(battleRecord)
	if err != nil {
		return nil, err
	}
	return createdBattleRecord, nil
}

// UC: 戦績編集 (パーティ戦績も更新)
func (s BattleRecordWriteService) EditBattleRecord(cmd command.EditBattleRecordCommand) (*battles.BattleRecord, error) {
	battleRecord, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	if _, err := s.spec.IsSatisfyForUpdate(battleRecord); err != nil {
		return nil, err
	}
	updatedBattleRecord, err := s.battleRecordRepo.Update(battleRecord)
	if err != nil {
		return nil, err
	}
	return updatedBattleRecord, nil
}

// UC: 戦績削除 (パーティ戦績も更新)
func (s BattleRecordWriteService) DeleteBattleRecord(cmd command.DeleteBattleRecord) (*battles.BattleRecord, error) {
	domain, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	if _, err := s.spec.ExistBattleRecord(domain.Id()); err != nil {
		return nil, err
	}
	return s.battleRecordRepo.Delete(domain.Id().Value())
}
