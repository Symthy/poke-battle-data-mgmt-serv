package items

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type IHeldItemNotification interface {
	SetId(identifier.HeldItemId)
	SetName(string)
	SetDescription(string)
	SetCorrectionValues(value.CorrectionValues)
}
