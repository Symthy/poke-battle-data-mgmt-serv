package battles

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type CorrectionApplier interface {
	Apply() uint16
}

type IPokemonBattleDataSet interface {
	PokemonBattleBaseParams
	TriggerConditionParams
}

type PokemonBattleBaseParams interface {
	IsNoDamage() bool
	MoveSpecies() value.MoveSpecies
	AttackPokemonActualValues() *value.PokemonActualValues
	AttackCorrectedActualValue() uint16
	PowerCorrectedValue() uint16
	MovePowerValue() uint16
	DefensePokemonActualValues() *value.PokemonActualValues
	DefenseCorrectedActualValue() uint16
	DamageCorrectedValue() uint16
	ApplyWeatherCorrection(damage uint16) uint16
	FieldCorrectedValue() uint16
	IsTypeMatch() bool
	TypeCompatibilityDamageRate() float32
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
