package abilities

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.AbilityId] = (*Ability)(nil)

type Ability struct {
	id               identifier.AbilityId
	name             string
	description      string
	correctionValues value.CorrectionValues
	// Todo
	triggerCondition *interface{}
}

func NewAbility(
	id identifier.AbilityId, name string, description string, correctionValues value.CorrectionValues,
) Ability {
	return Ability{
		id:               id,
		name:             name,
		description:      description,
		correctionValues: correctionValues,
	}
}

// Todo: refactor Notification
func (a Ability) Id() identifier.AbilityId {
	return a.id
}

func (a Ability) Notify(note IAbilityNotification) {
	note.SetId(a.id)
	note.SetName(a.name)
	note.SetDescription(a.description)
	note.SetCorrectionValues(a.correctionValues)
}
