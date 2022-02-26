package items

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.HeldItemId] = (*HeldItem)(nil)

type HeldItem struct {
	id               identifier.HeldItemId
	name             string
	description      string
	correctionValues []value.CorrectionValue
}

func NewHeldItem(
	id identifier.HeldItemId, name, description string, correctionValues []value.CorrectionValue,
) HeldItem {
	var corrections []value.CorrectionValue = nil
	if len(correctionValues) > 0 {
		corrections = correctionValues
	}
	return HeldItem{
		id:               id,
		name:             name,
		description:      description,
		correctionValues: corrections,
	}
}

func (i HeldItem) Id() identifier.HeldItemId {
	return i.id
}

func (i HeldItem) Notify(note IHeldItemNotification) {
	note.SetId(i.id)
	note.SetName(i.name)
	note.SetDescription(i.description)
	note.SetCorrectionValues(i.correctionValues)
}
