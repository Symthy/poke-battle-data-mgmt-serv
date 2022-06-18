package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type HeldItemInput struct {
	id            uint64
	name          string
	description   string
	battleEffects *battles.BattleEffects
}

func NewHeldItemInput(id uint64, name string, description string, battleEffects *battles.BattleEffects) HeldItemInput {
	return HeldItemInput{id, name, description, battleEffects}
}

func NewHeldItemBuilder() *HeldItemInput {
	return &HeldItemInput{}
}

func (b *HeldItemInput) Id(id uint64) {
	b.id = id
}
func (b *HeldItemInput) Name(name string) {
	b.name = name
}
func (b *HeldItemInput) Description(description string) {
	b.description = description
}
func (b *HeldItemInput) BattleEffects(battleEffects *battles.BattleEffects) {
	b.battleEffects = battleEffects
}

func (i HeldItemInput) BuildDomain() (*items.HeldItem, error) {
	id, err := identifier.NewHeldItemId(i.id)
	if err != nil {
		return nil, err
	}
	domain := items.NewHeldItem(*id, i.name, i.description, i.battleEffects)
	return &domain, nil
}
