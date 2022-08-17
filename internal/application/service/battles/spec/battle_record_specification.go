package spec

import (
	"fmt"

	"github.com/Symthy/PokeRest/internal/domain/entity/battles"
	"github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

// Todo: rename & split
type BattleRecordSpecification struct {
	battleRecordRepo repository.IBattleRecordRepository
	partyRepo        repository.IPartyRepository
}

func NewBattleRecordSpecification(
	battleRecordRepo repository.IBattleRecordRepository,
	partyRepo repository.IPartyRepository,
) BattleRecordSpecification {
	return BattleRecordSpecification{
		battleRecordRepo: battleRecordRepo,
		partyRepo:        partyRepo,
	}
}

func (s BattleRecordSpecification) IsSatisfyForUpdate(battleRecord *battles.BattleRecord) (bool, error) {
	if _, err := s.ExistSelfParty(battleRecord.PartyId()); err != nil {
		return false, err
	}
	if _, err := s.ExistBattleRecord(battleRecord.Id()); err != nil {
		return false, err
	}
	return true, nil
}

func (s BattleRecordSpecification) ExistSelfParty(partyId identifier.PartyId) (bool, error) {
	party, err := s.partyRepo.FindById(partyId.Value())
	if err != nil {
		return false, err
	}
	if party == nil {
		// Todo: replace error
		return false, fmt.Errorf("target party not found")
	}
	return true, nil
}

func (s BattleRecordSpecification) ExistBattleRecord(id identifier.BattleRecordId) (bool, error) {
	retBattleRecord, err := s.battleRecordRepo.FindById(id.Value())
	if err != nil {
		return false, err
	}
	if retBattleRecord == nil {
		// Todo: replace error
		return false, fmt.Errorf("target battle record not found")
	}
	return true, nil
}
