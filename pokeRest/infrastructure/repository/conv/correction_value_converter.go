package conv

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

func toCorrectionValues(correction mixin.CorrectionValue) []value.CorrectionValue {
	values := []value.CorrectionValue{}
	if correction.AttackCorrectionValue != 1 {
		value := value.AttackCorrectionValue(nil, correction.AttackCorrectionValue)
		values = append(values, *value)
	}
	if correction.SpecialAttackCorrectionValue != 1 {
		value := value.SpecialAttackCorrectionValue(nil, correction.SpecialAttackCorrectionValue)
		values = append(values, *value)
	}
	if correction.AttackPowerCorrectionValue != 1 {
		value := value.AttackPowerCorrectionValue(nil, correction.AttackPowerCorrectionValue)
		values = append(values, *value)
	}
	if correction.SpecialAttackPowerCorrectionValue != 1 {
		value := value.SpecialAttackPowerCorrectionValue(nil, correction.SpecialAttackPowerCorrectionValue)
		values = append(values, *value)
	}
	if correction.PhysicalMovePowerCorrectionValue != 1 {
		value := value.PhysicalMoveCorrectionValue(nil, correction.PhysicalMovePowerCorrectionValue)
		values = append(values, *value)
	}
	if correction.SpecialMovePowerCorrectionValue != 1 {
		value := value.SpecialMoveCorrectionValue(nil, correction.SpecialMovePowerCorrectionValue)
		values = append(values, *value)
	}
	if correction.DeffenceCorrectionValue != 1 {
		value := value.DeffenceCorrectionValue(nil, correction.DeffenceCorrectionValue)
		values = append(values, *value)
	}
	if correction.SpecialDeffenceCorrectionValue != 1 {
		value := value.SpecialDeffenceCorrectionValue(nil, correction.SpecialDeffenceCorrectionValue)
		values = append(values, *value)
	}
	if correction.DamageCorrectionValue1 != 1 {
		t := value.NewPokemonType(correction.DamageCorrectionType1.String())
		value := value.DamageCorrectionValue(t, correction.DamageCorrectionValue1)
		values = append(values, *value)
	}
	if correction.DamageCorrectionValue2 != 1 {
		t := value.NewPokemonType(correction.DamageCorrectionType2.String())
		value := value.DamageCorrectionValue(t, correction.DamageCorrectionValue2)
		values = append(values, *value)
	}
	return values
}
