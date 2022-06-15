package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/effects"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type HeldItemInput struct {
	id            uint64
	name          string
	description   string
	battleEffects *effects.BattleEffects
}

func NewHeldItemInput(id uint64, name string, description string, battleEffects *effects.BattleEffects) HeldItemInput {
	return HeldItemInput{id, name, description, battleEffects}
}

func (i HeldItemInput) BuildDomain() (*items.HeldItem, error) {
	id, err := identifier.NewHeldItemId(i.id)
	if err != nil {
		return nil, err
	}
	domain := items.NewHeldItem(*id, i.name, i.description, i.battleEffects)
	return &domain, nil
}
