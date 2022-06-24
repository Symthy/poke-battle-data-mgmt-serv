package battles

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type IPokemonBattleDataSet interface {
	PokemonBattleBaseParam
	TriggerConditionParams
}

type PokemonBattleBaseParam interface {
	IsNoDamage() bool
	MoveSpecies() value.MoveSpecies
	AttackPokemonActualValues() *value.PokemonActualValues
	AttackPokemonCorrectedActualValueA() uint16
	AttackPokemonCorrectedActualValueC() uint16
	AttackPowerValue() uint16
	MovePowerValue() uint16
	DefensePokemonActualValues() *value.PokemonActualValues
	DefensePokemonCorrectedActualValueB() uint16
	DefensePokemonCorrectedActualValueD() uint16
}

type TriggerConditionParams interface {
	AttackPokemonTypeOfFirst() value.PokemonType
	AttackPokemonTypeOfSecond() value.PokemonType
	AttackPokemonCorrectedActualValueS() uint16
	DefensePokemonTypeOfFirst() value.PokemonType
	DefensePokemonTypeOfSecond() value.PokemonType
	DefensePokemonCorrectedActualValueS() uint16
	MovePokemonType() value.PokemonType
	HasItemAttackSide() bool
	HasItemDefenseSide() bool
	TypeCompatibilityDamageRate() float32
}
