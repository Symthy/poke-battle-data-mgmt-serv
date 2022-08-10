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
	WeatherCorrectedValue() uint16
	FieldCorrectedValue() uint16
	IsTypeMatchAttackSide() bool
	IsBurnAttackSide() bool
	AbnormalStateAttackSideCorectedValue() uint16
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
}

var _ IPokemonBattleDataSet = (*EmptyPokemonBattleDataSet)(nil)

type EmptyPokemonBattleDataSet struct {
	PokemonBattleBaseParams
	EmptyTriggerConditionParams
}

var _ PokemonBattleBaseParams = (*EmptyPokemonBattleBaseParams)(nil)

type EmptyPokemonBattleBaseParams struct {
}

func (p EmptyPokemonBattleBaseParams) IsNoDamage() bool {
	return false
}
func (p EmptyPokemonBattleBaseParams) MoveSpecies() value.MoveSpecies {
	return value.MoveSpeciesUnknown
}
func (p EmptyPokemonBattleBaseParams) AttackPokemonActualValues() *value.PokemonActualValues {
	return nil
}
func (p EmptyPokemonBattleBaseParams) AttackCorrectedActualValue() uint16 {
	return 0
}
func (p EmptyPokemonBattleBaseParams) PowerCorrectedValue() uint16 {
	return 0
}
func (p EmptyPokemonBattleBaseParams) MovePowerValue() uint16 {
	return 0
}
func (p EmptyPokemonBattleBaseParams) DefensePokemonActualValues() *value.PokemonActualValues {
	return nil
}
func (p EmptyPokemonBattleBaseParams) DefenseCorrectedActualValue() uint16 {
	return 0
}
func (p EmptyPokemonBattleBaseParams) DamageCorrectedValue() uint16 {
	return 0
}
func (p EmptyPokemonBattleBaseParams) WeatherCorrectedValue() uint16 {
	return 0
}
func (p EmptyPokemonBattleBaseParams) FieldCorrectedValue() uint16 {
	return 0
}
func (p EmptyPokemonBattleBaseParams) IsTypeMatchAttackSide() bool {
	return false
}
func (p EmptyPokemonBattleBaseParams) IsBurnAttackSide() bool {
	return false
}
func (p EmptyPokemonBattleBaseParams) AbnormalStateAttackSideCorectedValue() uint16 {
	return 0
}
func (p EmptyPokemonBattleBaseParams) TypeCompatibilityDamageRate() float32 {
	return 1.0
}

var _ TriggerConditionParams = (*EmptyTriggerConditionParams)(nil)

type EmptyTriggerConditionParams struct {
}

func (m EmptyTriggerConditionParams) AttackPokemonTypeOfFirst() value.PokemonType {
	return value.UnknownType()
}
func (m EmptyTriggerConditionParams) AttackPokemonTypeOfSecond() value.PokemonType {
	return value.UnknownType()
}
func (m EmptyTriggerConditionParams) AttackPokemonCorrectedActualValueS() uint16 {
	return 0
}
func (m EmptyTriggerConditionParams) DefensePokemonTypeOfFirst() value.PokemonType {
	return value.UnknownType()
}
func (m EmptyTriggerConditionParams) DefensePokemonTypeOfSecond() value.PokemonType {
	return value.UnknownType()
}
func (m EmptyTriggerConditionParams) DefensePokemonCorrectedActualValueS() uint16 {
	return 0
}
func (m EmptyTriggerConditionParams) MovePokemonType() value.PokemonType {
	return value.UnknownType()
}
func (m EmptyTriggerConditionParams) HasItemAttackSide() bool {
	return false
}
func (m EmptyTriggerConditionParams) HasItemDefenseSide() bool {
	return false
}
