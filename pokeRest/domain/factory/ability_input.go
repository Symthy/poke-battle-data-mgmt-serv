package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type AbilityInput struct {
	id            uint
	name          string
	description   string
	battleEffects *value.BattleEffects
}

func NewAbilityInput(id uint, name string, description string, battleEffects *value.BattleEffects) AbilityInput {
	return AbilityInput{
		id:            id,
		name:          name,
		description:   description,
		battleEffects: battleEffects,
	}
}

func (a AbilityInput) BuildDomain() (*abilities.Ability, error) {
	id, err := identifier.NewAbilityId(a.id)
	if err != nil {
		return nil, err
	}
	domain := abilities.NewAbility(*id, a.name, a.description, a.battleEffects)
	return &domain, nil
}
