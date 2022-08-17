package items

import (
	"github.com/Symthy/PokeRest/internal/domain/value/battles"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type IHeldItemNotification interface {
	SetId(identifier.HeldItemId)
	SetName(string)
	SetDescription(string)
	SetBattleEffects(*battles.BattleEffects)
}
