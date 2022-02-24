package spec

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

// Todo: rename & split
type BattleRecordSpecification struct {
	seasonRepo        repository.IBattleSeasonRepository
	partyRepo         repository.IPartyRepository
	opponentPartyRepo repository.IBattleOpponentPartyRepository
}

func NewBattleRecordSpecification(
	seasonRepo repository.IBattleSeasonRepository,
	partyRepo repository.IPartyRepository,
	opponentPartyRepo repository.IBattleOpponentPartyRepository,
) BattleRecordSpecification {
	return BattleRecordSpecification{
		seasonRepo:        seasonRepo,
		partyRepo:         partyRepo,
		opponentPartyRepo: opponentPartyRepo,
	}
}

func (b BattleRecordSpecification) ResolveCurrentSeason() (*battles.Season, error) {
	currentSeason, err := b.seasonRepo.FindCurrent()
	if err != nil {
		return nil, err
	}
	return currentSeason, nil
}

func (b BattleRecordSpecification) ExistSelfParty(partyId uint) (bool, error) {
	party, err := b.partyRepo.FindById(partyId)
	if party == nil {
		return false, err
	}
	return true, nil
}

func (b BattleRecordSpecification) ResolveOppositeParty(
	opponentPartyMember value.PartyPokemonIds) (*battles.BattleOpponentParty, error) {
	opponentParty, err := b.opponentPartyRepo.FindParty(opponentPartyMember)
	if err != nil {
		return nil, err
	}
	if opponentParty == nil {
		created, err := b.opponentPartyRepo.Create(battles.NewBattleOpponentPartyOfUnregister(opponentPartyMember))
		if err != nil {
			return nil, err
		}
		opponentParty = created
	}
	return opponentParty, nil
}
