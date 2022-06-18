package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type AbilityInput struct {
	id            uint64
	name          string
	description   string
	battleEffects *battles.BattleEffects
}

func NewAbilityInput(id uint64, name string, description string, battleEffects *battles.BattleEffects) AbilityInput {
	return AbilityInput{
		id:            id,
		name:          name,
		description:   description,
		battleEffects: battleEffects,
	}
}

func NewAbilityBuilder() *AbilityInput {
	return &AbilityInput{}
}

func (b *AbilityInput) Id(id uint64) {
	b.id = id
}
func (b *AbilityInput) Name(name string) {
	b.name = name
}
func (b *AbilityInput) Description(description string) {
	b.description = description
}
func (b *AbilityInput) BattleEffects(battleEffects *battles.BattleEffects) {
	b.battleEffects = battleEffects
}

func (a AbilityInput) BuildDomain() (*abilities.Ability, error) {
	id, err := identifier.NewAbilityId(a.id)
	if err != nil {
		return nil, err
	}
	domain := abilities.NewAbility(*id, a.name, a.description, a.battleEffects)
	return &domain, nil
}
