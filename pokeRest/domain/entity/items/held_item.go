package items

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

var _ entity.IDomain = (*HeldItem)(nil)

type HeldItem struct {
	id               uint
	name             string
	description      string
	correctionValues *[]value.CorrectionValue
}

func NewHeldItem(id uint, name, description string, correctionValues []value.CorrectionValue) HeldItem {
	var corrections *[]value.CorrectionValue = nil
	if len(correctionValues) > 0 {
		corrections = &correctionValues
	}
	return HeldItem{
		id:               id,
		name:             name,
		description:      description,
		correctionValues: corrections,
	}
}

func (i HeldItem) Id() uint {
	return i.id
}

func (a HeldItem) Name() string {
	return a.name
}

func (a HeldItem) Description() string {
	return a.description
}

func (a HeldItem) CorrectionValues() *[]value.CorrectionValue {
	return a.correctionValues
}
