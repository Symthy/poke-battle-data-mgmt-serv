package abilities

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type IAbilityNotification interface {
	SetId(identifier.AbilityId)
	SetName(string)
	SetDescription(string)
	SetCorrectionValues([]value.CorrectionValue)
}
