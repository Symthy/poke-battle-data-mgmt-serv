package service

import (
	s_damages "github.com/Symthy/PokeRest/pokeRest/application/service/damages"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type IBattleEffectsResolver interface {
	resolve(s_damages.BattlePokemonAdjustment) (*value.BattleEffects, error)
}

type BattleEffectsResolver struct {
	abilityRepo  repository.IAbilityRepository
	heldItemRepo repository.IHeldItemRepository
}

func (r BattleEffectsResolver) resolveBattleEffects(
	adjustment s_damages.BattlePokemonAdjustment) (*value.BattleEffects, error) {
	effects := &value.BattleEffects{}
	ability, err := r.abilityRepo.FindById(adjustment.AbilityId())
	if err != nil {
		return nil, err
	}
	ability.NotifyBattleEffects(effects)
	item, err := r.heldItemRepo.FindById(adjustment.HeldItemId())
	if err != nil {
		return nil, err
	}
	item.NotifyBattleEffects(effects)
	return effects, err
}
