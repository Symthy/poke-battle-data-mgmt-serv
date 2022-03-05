package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type AbilityInput struct {
	id               uint
	name             string
	description      string
	correctionValues []value.CorrectionValue
}

func NewAbilityInput(id uint, name string, description string, correctionValues []value.CorrectionValue) AbilityInput {
	return AbilityInput{
		id:               id,
		name:             name,
		description:      description,
		correctionValues: correctionValues,
	}
}

func (a AbilityInput) BuildDomain() (*abilities.Ability, error) {
	id, err := identifier.NewAbilityId(a.id)
	if err != nil {
		return nil, err
	}
	correctionValues := value.NewCorrectionValues(a.correctionValues)
	domain := abilities.NewAbility(*id, a.name, a.description, correctionValues)
	return &domain, nil
}
