package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type HeldItemInput struct {
	id               uint
	name             string
	description      string
	correctionValues []value.CorrectionValue
}

func NewHeldItemInput(id uint, name string, description string, correctionValues []value.CorrectionValue) HeldItemInput {
	return HeldItemInput{id, name, description, correctionValues}
}

func (i HeldItemInput) BuildDomain() (*items.HeldItem, error) {
	id, err := identifier.NewHeldItemId(i.id)
	if err != nil {
		return nil, err
	}
	values := value.NewCorrectionValues(i.correctionValues)
	domain := items.NewHeldItem(*id, i.name, i.description, values)
	return &domain, nil
}
