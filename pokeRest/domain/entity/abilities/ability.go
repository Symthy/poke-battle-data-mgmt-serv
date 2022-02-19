package abilities

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type Ability struct {
	id               uint
	name             string
	description      string
	correctionValues *[]value.CorrectionValue
	// Todo
	triggerCondition *interface{}
}

func NewAbility(id uint, name string, description string, correctionValues *[]value.CorrectionValue) Ability {
	return Ability{
		id:               id,
		name:             name,
		description:      description,
		correctionValues: correctionValues,
	}
}

func (a Ability) Id() uint {
	return a.id
}

func (a Ability) Name() string {
	return a.name
}

func (a Ability) Description() string {
	return a.description
}

func (a Ability) CorrectionValues() *[]value.CorrectionValue {
	return a.correctionValues
}
