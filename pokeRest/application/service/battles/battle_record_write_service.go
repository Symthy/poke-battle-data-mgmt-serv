package battles

import (
	"fmt"

	"github.com/Symthy/PokeRest/pokeRest/application/service/battles/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	d_service "github.com/Symthy/PokeRest/pokeRest/domain/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type battleRecordtransactionalRepositoryBuilder = func(repository.IBattleRecordRepository) repository.IBattleRecordTransactionalRepository

type BattleRecordWriteService struct {
	// Todo: 個別に分けるかどうか
	battleRecordRepo               repository.IBattleRecordRepository
	opponentPartyRepo              repository.IBattleOpponentPartyRepository
	selfPartyRepo                  repository.IPartyRepository
	seasonRepo                     repository.IBattleSeasonRepository
	battleRecordResolver           d_service.BattleRecordResolver
	transactionalRepositoryBuilder battleRecordtransactionalRepositoryBuilder
}

func NewBattleRecordWriteService(
	battleRecordRepo repository.IBattleRecordRepository,
	opponentPartyRepo repository.IBattleOpponentPartyRepository,
	selfPartyRepo repository.IPartyRepository,
	seasonRepo repository.IBattleSeasonRepository,
	transactionalRepositoryBuilder battleRecordtransactionalRepositoryBuilder,
) BattleRecordWriteService {
	return BattleRecordWriteService{
		battleRecordRepo:               battleRecordRepo,
		opponentPartyRepo:              opponentPartyRepo,
		selfPartyRepo:                  selfPartyRepo,
		seasonRepo:                     seasonRepo,
		battleRecordResolver:           d_service.NewBattleRecordResolver(opponentPartyRepo, selfPartyRepo, seasonRepo),
		transactionalRepositoryBuilder: transactionalRepositoryBuilder,
	}
}

// UC: 戦績登録 (パーティ戦績も更新)
func (s BattleRecordWriteService) AddBattleRecord(cmd command.AddBattleRecordCommand) (*battles.BattleRecord, error) {
	input, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	opponentPartyMember := value.NewPartyPokemonIds(cmd.OpponentPartyPokemonIds())
	transactionalRepo := s.transactionalRepositoryBuilder(s.battleRecordRepo)

	if err := transactionalRepo.StartTransaction(); err != nil {
		return nil, err
	}
	defer transactionalRepo.PanicPostProcess()

	battleRecord, err := s.battleRecordResolver.Resolve(*input, opponentPartyMember)
	if err != nil {
		transactionalRepo.CancelTransaction()
		return nil, err
	}

	createdBattleRecord, err := transactionalRepo.Create(*battleRecord)
	if err != nil {
		transactionalRepo.CancelTransaction()
		return nil, err
	}

	if err := transactionalRepo.FinishTransaction(); err != nil {
		return nil, err
	}
	return createdBattleRecord, nil
}

// UC: 戦績編集 (パーティ戦績も更新)
func (s BattleRecordWriteService) EditBattleRecord(cmd command.EditBattleRecordCommand) (*battles.BattleRecord, error) {
	input, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	opponentPartyMember := value.NewPartyPokemonIds(cmd.OpponentPartyPokemonIds())
	transactionalRepo := s.transactionalRepositoryBuilder(s.battleRecordRepo)

	if err := transactionalRepo.StartTransaction(); err != nil {
		return nil, err
	}
	defer transactionalRepo.PanicPostProcess()

	if err := s.validateBattleRecord(*input); err != nil {
		return nil, err
	}
	battleRecord, err := s.battleRecordResolver.Resolve(*input, opponentPartyMember)
	if err != nil {
		transactionalRepo.CancelTransaction()
		return nil, err
	}
	updatedBattleRecord, err := transactionalRepo.Update(*battleRecord)
	if err != nil {
		transactionalRepo.CancelTransaction()
		return nil, err
	}

	if err := transactionalRepo.FinishTransaction(); err != nil {
		return nil, err
	}
	return updatedBattleRecord, nil
}

// UC: 戦績削除 (パーティ戦績も更新)
func (s BattleRecordWriteService) DeleteBattleRecord(id uint) (*battles.BattleRecord, error) {
	return s.battleRecordRepo.Delete(id)
}

func (s BattleRecordWriteService) validateBattleRecord(battleRecord battles.BattleRecord) error {
	retBattleRecord, err := s.battleRecordRepo.FindById(battleRecord.Id().Value())
	if err != nil {
		return err
	}
	if retBattleRecord == nil {
		// Todo: replace error
		return fmt.Errorf("target battle record not found")
	}
	return nil
}
