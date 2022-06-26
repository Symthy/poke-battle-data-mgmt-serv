package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/common/lists"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

type FieldStateType string

const (
	ElectricField FieldStateType = "Electric Terrain"
	GrassyField   FieldStateType = "Grassy Terrain"
	MistyField    FieldStateType = "Misty Terrain"
	PsychicField  FieldStateType = "Psychic Terrain"
	NormalField   FieldStateType = ""
)

func allFieldStateType() []FieldStateType {
	return []FieldStateType{ElectricField, GrassyField, MistyField, PsychicField}
}

type FieldState struct {
	field       FieldStateType
	corrections *battles.BattleCorrectionValues
}

func NewFieldState(field string) FieldState {
	if !lists.Contains(allFieldStateType(), field) {
		return FieldState{
			field:       FieldStateType(field),
			corrections: battles.NewBattleCorrectionValues(),
		}
	}
	return FieldState{
		field:       FieldStateType(field),
		corrections: resolveFieldCorrection(FieldStateType(field)),
	}
}

func (f FieldState) ApplyCorrection(data battles.IPokemonBattleDataSet) uint16 {
	if f.field == NormalField {
		return 4096
	}
	target := battles.CorrectionNone
	if data.MoveSpecies().IsPhysical() {
		target = battles.CorrectionPhysicalPower
	}
	if data.MoveSpecies().IsSpecial() {
		target = battles.CorrectionSpecialPower
	}
	return f.corrections.Apply(4096, target, data, battles.BattleAttackSide)
}

func resolveFieldCorrection(field FieldStateType) *battles.BattleCorrectionValues {
	if field == ElectricField {
		return battles.NewBattleCorrectionValues(
			battles.NewBattleCorrectionValue(
				battles.CorrectionPhysicalPower,
				5325,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Electric().ToString())),
			battles.NewBattleCorrectionValue(
				battles.CorrectionSpecialPower,
				5325,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Electric().ToString())),
		)
	}
	if field == GrassyField {
		return battles.NewBattleCorrectionValues(
			battles.NewBattleCorrectionValue(
				battles.CorrectionPhysicalPower,
				5325,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Grass().ToString())),
			battles.NewBattleCorrectionValue(
				battles.CorrectionSpecialPower,
				5325,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Grass().ToString())),
		)
	}
	if field == MistyField {
		return battles.NewBattleCorrectionValues(
			battles.NewBattleCorrectionValue(
				battles.CorrectionPhysicalPower,
				2048,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Dragon().ToString())),
			battles.NewBattleCorrectionValue(
				battles.CorrectionSpecialPower,
				2048,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Dragon().ToString())),
		)
	}
	if field == PsychicField {
		return battles.NewBattleCorrectionValues(
			battles.NewBattleCorrectionValue(
				battles.CorrectionPhysicalPower,
				5325,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Psychic().ToString())),
			battles.NewBattleCorrectionValue(
				battles.CorrectionSpecialPower,
				5325,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Psychic().ToString())),
		)
	}
	return battles.NewBattleCorrectionValues()
}
