package service

import (
	"fmt"

	"github.com/Symthy/PokeRest/pokeRest/application/service/battles/spec"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type BattleRecordResolver struct {
	battleRecordSpec spec.BattleRecordSpecification
}

func NewBattleRecordResolver(
	opponentPartyRepo repository.IBattleOpponentPartyRepository,
	selfPartyRepo repository.IPartyRepository,
	seasonRepo repository.IBattleSeasonRepository) BattleRecordResolver {
	return BattleRecordResolver{
		battleRecordSpec: spec.NewBattleRecordSpecification(seasonRepo, selfPartyRepo, opponentPartyRepo),
	}
}

func (r BattleRecordResolver) Resolve(
	input battles.BattleRecord, opponentPartyMember value.PartyPokemonIds,
) (*battles.BattleRecord, error) {
	battleRecord := input
	if err := r.validateSelfParty(input); err != nil {
		return nil, err
	}
	currentSeason, err := r.battleRecordSpec.ResolveCurrentSeason()
	if err != nil {
		return nil, err
	}
	battleRecord.ApplySeason(*currentSeason)

	opponentParty, err := r.battleRecordSpec.ResolveOppositeParty(opponentPartyMember)
	if err != nil {
		return nil, err
	}
	battleRecord.ApplyOpponentPartyId(opponentParty.Id())
	return &battleRecord, nil
}

func (r BattleRecordResolver) validateSelfParty(battleRecord battles.BattleRecord) error {
	isExistParty, err := r.battleRecordSpec.ExistSelfParty(battleRecord.PartyId())
	if err != nil {
		return err
	}
	if !isExistParty {
		// Todo: replace error
		return fmt.Errorf("self party not found")
	}
	return nil
}
