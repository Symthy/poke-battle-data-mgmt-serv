package moves

import (
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type IMoveNotification interface {
	SetId(identifier.MoveId)
	SetName(string)
	SetSpecies(value.MoveSpecies)
	SetPower(uint16)         // 威力
	SetAccuracyRate(float32) // 命中率
	SetPP(uint8)
	SetIsContained(bool)
	SetCanGuard(bool)
}

type IMoveNotificationForDamageCalc interface {
	SetMoveType(value.PokemonType)
	SetMoveSpecies(value.MoveSpecies)
	SetMovePower(uint16)
	SetIsMoveContained(bool)
	SetCanMoveGuard(bool)
}
