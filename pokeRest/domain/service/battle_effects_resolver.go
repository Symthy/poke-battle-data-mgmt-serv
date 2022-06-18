package service

import (
	s_damages "github.com/Symthy/PokeRest/pokeRest/application/service/damages"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

type BattleEffectsResolver struct {
	abilityRepo  repository.IAbilityRepository
	heldItemRepo repository.IHeldItemRepository
}

func NewBattleEffectsResolver(abilityRepo repository.IAbilityRepository, heldItemRepo repository.IHeldItemRepository) *BattleEffectsResolver {
	return &BattleEffectsResolver{
		abilityRepo:  abilityRepo,
		heldItemRepo: heldItemRepo,
	}
}

func (r BattleEffectsResolver) resolve(adjustment *s_damages.BattlePokemonAdjustment) (*battles.BattleEffects, error) {
	effects := battles.NewEmptyBattleEffects()
	ability, err := r.abilityRepo.FindById(adjustment.AbilityId().Value())
	if err != nil {
		return nil, err
	}
	ability.NotifyBattleEffects(effects)
	item, err := r.heldItemRepo.FindById(adjustment.HeldItemId().Value())
	if err != nil {
		return nil, err
	}
	item.NotifyBattleEffects(effects)
	return effects, nil
}
