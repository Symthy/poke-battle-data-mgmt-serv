package items

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type IHeldItemNotification interface {
	SetId(identifier.HeldItemId)
	SetName(string)
	SetDescription(string)
	SetBattleEffects(*battles.BattleEffects)
}
