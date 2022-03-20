package moves

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type IMoveNotification interface {
	SetId(identifier.MoveId)
	SetName(string)
	SetSpecies(value.MoveSpecies)
	SetPower(int)            // 威力
	SetAccuracyRate(float32) // 命中率
	SetPP(int)
	SetIsContained(bool)
	SetCanGuard(bool)
}

type IMoveNotificationForDamageCalc interface {
	SetMoveType(value.PokemonType)
	SetMoveSpecies(value.MoveSpecies)
	SetMovePower(int)
	SetIsMoveContained(bool)
	SetCanMoveGuard(bool)
}
