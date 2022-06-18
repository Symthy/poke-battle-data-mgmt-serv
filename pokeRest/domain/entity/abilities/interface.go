package abilities

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type IAbilityNotification interface {
	SetId(identifier.AbilityId)
	SetName(string)
	SetDescription(string)
	SetBattleEffects(*battles.BattleEffects)
}
