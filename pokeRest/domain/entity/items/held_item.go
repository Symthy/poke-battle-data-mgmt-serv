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

func NewHeldItem(id uint) HeldItem {
	return HeldItem{
		id: id,
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
