package abilities

import (
	"github.com/Symthy/PokeRest/internal/domain/value/battles"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type IAbilityNotification interface {
	SetId(identifier.AbilityId)
	SetName(string)
	SetDescription(string)
	SetBattleEffects(*battles.BattleEffects)
}
