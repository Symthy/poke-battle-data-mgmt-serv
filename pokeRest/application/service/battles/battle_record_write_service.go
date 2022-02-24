package battles

import (
	"fmt"

	"github.com/Symthy/PokeRest/pokeRest/application/service/battles/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	d_service "github.com/Symthy/PokeRest/pokeRest/domain/service"
)

type BattleRecordWriteService struct {
	// Todo: 個別に分けるかどうか
	battleRecordRepo     repository.IBattleRecordTransactionalRepository
	opponentPartyRepo    repository.IBattleOpponentPartyRepository
	selfPartyRepo        repository.IPartyRepository
	seasonRepo           repository.IBattleSeasonRepository
	battleRecordResolver d_service.BattleRecordResolver
}

func NewBattleRecordWriteService(
	battleRecordRepo repository.IBattleRecordTransactionalRepository,
	opponentPartyRepo repository.IBattleOpponentPartyRepository,
	selfPartyRepo repository.IPartyRepository,
	seasonRepo repository.IBattleSeasonRepository) BattleRecordWriteService {
	return BattleRecordWriteService{
		battleRecordRepo:     battleRecordRepo,
		opponentPartyRepo:    opponentPartyRepo,
		selfPartyRepo:        selfPartyRepo,
		seasonRepo:           seasonRepo,
		battleRecordResolver: d_service.NewBattleRecordResolver(opponentPartyRepo, selfPartyRepo, seasonRepo),
	}
}

// UC: 戦績登録 (パーティ戦績も更新)
func (s BattleRecordWriteService) AddBattleRecord(cmd command.AddBattleRecordCommand) (*battles.BattleRecord, error) {
	inputBattleRecord := cmd.ToDomain()
	opponentPartyMember := cmd.OpponentPartyMember()

	if err := s.battleRecordRepo.StartTransaction(); err != nil {
		return nil, err
	}
	defer s.battleRecordRepo.PanicPostProcess()

	battleRecord, err := s.battleRecordResolver.Resolve(inputBattleRecord, opponentPartyMember)
	if err != nil {
		s.battleRecordRepo.CancelTransaction()
		return nil, err
	}

	createdBattleRecord, err := s.battleRecordRepo.Create(*battleRecord)
	if err != nil {
		s.battleRecordRepo.CancelTransaction()
		return nil, err
	}

	if err := s.battleRecordRepo.FinishTransaction(); err != nil {
		return nil, err
	}
	return createdBattleRecord, nil
}

// UC: 戦績編集 (パーティ戦績も更新)
func (s BattleRecordWriteService) EditBattleRecord(cmd command.EditBattleRecordCommand) (*battles.BattleRecord, error) {
	inputBattleRecord := cmd.ToDomain()
	opponentPartyMember := cmd.OpponentPartyMember()

	if err := s.battleRecordRepo.StartTransaction(); err != nil {
		return nil, err
	}
	defer s.battleRecordRepo.PanicPostProcess()

	if err := s.validateBattleRecord(inputBattleRecord); err != nil {
		return nil, err
	}
	battleRecord, err := s.battleRecordResolver.Resolve(inputBattleRecord, opponentPartyMember)
	if err != nil {
		s.battleRecordRepo.CancelTransaction()
		return nil, err
	}
	updatedBattleRecord, err := s.battleRecordRepo.Update(*battleRecord)
	if err != nil {
		s.battleRecordRepo.CancelTransaction()
		return nil, err
	}

	if err := s.battleRecordRepo.FinishTransaction(); err != nil {
		return nil, err
	}
	return updatedBattleRecord, nil
}

// UC: 戦績削除 (パーティ戦績も更新)
func (s BattleRecordWriteService) DeleteBattleRecord(id uint) (*battles.BattleRecord, error) {
	return s.battleRecordRepo.Delete(id)
}

func (s BattleRecordWriteService) validateBattleRecord(battleRecord battles.BattleRecord) error {
	retBattleRecord, err := s.battleRecordRepo.FindById(battleRecord.Id())
	if err != nil {
		return err
	}
	if retBattleRecord == nil {
		// Todo: replace error
		return fmt.Errorf("target battle record not found")
	}
	return nil
}
