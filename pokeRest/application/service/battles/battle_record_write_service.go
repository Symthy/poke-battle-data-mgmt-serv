package battles

import (
	"fmt"

	"github.com/Symthy/PokeRest/pokeRest/application/service/battles/command"
	"github.com/Symthy/PokeRest/pokeRest/application/service/battles/spec"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type BattleRecordWriteServie struct {
	// Todo: 個別に分けるかどうか
	battleRecordRepo  repository.IBattleRecordTransactionalRepository
	opponentPartyRepo repository.IBattleOpponentPartyRepository
	selfPartyRepo     repository.IPartyRepository
	seasonRepo        repository.IBattleSeasonRepository
	battleRecordSpec  spec.BattleRecordSpecification
}

func NewBattleRecordWriteServie(
	battleRecordRepo repository.IBattleRecordTransactionalRepository,
	opponentPartyRepo repository.IBattleOpponentPartyRepository,
	selfPartyRepo repository.IPartyRepository,
	seasonRepo repository.IBattleSeasonRepository) BattleRecordWriteServie {
	return BattleRecordWriteServie{
		battleRecordRepo:  battleRecordRepo,
		opponentPartyRepo: opponentPartyRepo,
		selfPartyRepo:     selfPartyRepo,
		seasonRepo:        seasonRepo,
		battleRecordSpec:  spec.NewBattleSpecification(seasonRepo, selfPartyRepo, opponentPartyRepo),
	}
}

// UC: 戦績登録 (パーティ戦績も更新)
func (s BattleRecordWriteServie) AddBattleRecord(cmd command.AddBattleRecordCommand) (*battles.BattleRecord, error) {
	// Todo: split domain service
	battleRecord := cmd.ToDomain()

	if err := s.battleRecordRepo.StartTransaction(); err != nil {
		return nil, err
	}
	defer s.battleRecordRepo.PanicPostProcess()

	s.validateSelfParty(battleRecord)
	currentSeason, err := s.battleRecordSpec.ResolveCurrentSeason()
	if err != nil {
		s.battleRecordRepo.CancelTransaction()
		return nil, err
	}
	battleRecord.ApplySeason(*currentSeason)

	opponentParty, err := s.battleRecordSpec.ResolveOppositeParty(cmd.OpponentPartyMember())
	if err != nil {
		s.battleRecordRepo.CancelTransaction()
		return nil, err
	}
	battleRecord.ApplyOpponentPartyId(opponentParty.Id())

	createdBattleRecord, err := s.battleRecordRepo.Create(battleRecord)
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
func (s BattleRecordWriteServie) EditBattleRecord() (*battles.BattleRecord, error) {
	return nil, nil
}

// UC: 戦績削除 (パーティ戦績も更新)
func (s BattleRecordWriteServie) DeleteBattleRecord() (*battles.BattleRecord, error) {
	return nil, nil
}

func (s BattleRecordWriteServie) validateSelfParty(battleRecord battles.BattleRecord) error {
	isExistParty, err := s.battleRecordSpec.ExistSelfParty(battleRecord.PartyId())
	if err != nil {
		return err
	}
	if !isExistParty {
		// Todo: replace error
		return fmt.Errorf("self party not found")
	}
	return nil
}
