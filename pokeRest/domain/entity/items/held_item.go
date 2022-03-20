package items

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.HeldItemId] = (*HeldItem)(nil)

type HeldItem struct {
	id            identifier.HeldItemId
	name          string
	description   string
	battleEffects value.BattleEffects
}

func NewHeldItem(
	id identifier.HeldItemId, name, description string, battleEffects value.BattleEffects,
) HeldItem {
	return HeldItem{
		id:            id,
		name:          name,
		description:   description,
		battleEffects: battleEffects,
	}
}

func (i HeldItem) Id() identifier.HeldItemId {
	return i.id
}

func (i HeldItem) NotifyBattleEffects(effects value.BattleEffects) {
	effects.Merge(i.battleEffects)
}

func (i HeldItem) Notify(note IHeldItemNotification) {
	note.SetId(i.id)
	note.SetName(i.name)
	note.SetDescription(i.description)
	note.SetBattleEffects(i.battleEffects)
}
