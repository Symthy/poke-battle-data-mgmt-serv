package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

type WeatherState struct {
	weather     WeatherType
	corrections *battles.BattleCorrectionValues
}

func NewWeatherState(weather string) WeatherState {
	return WeatherState{
		weather:     WeatherType(weather),
		corrections: resolveWeatherCorrection(WeatherType(weather)),
	}
}

func (w WeatherState) ApplyCorrection(damage uint16, data battles.IPokemonBattleDataSet) uint16 {
	if w.weather == WeatherNormal {
		return damage
	}
	return w.corrections.Apply(damage, battles.CorrectionDamage, data, battles.BattleAttackSide)
}

func resolveWeatherCorrection(weather WeatherType) *battles.BattleCorrectionValues {
	if weather == WeatherSunny {
		return battles.NewBattleCorrectionValues(
			battles.NewBattleCorrectionValue(
				battles.CorrectionDamage,
				6144,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Fire().ToString())),
			battles.NewBattleCorrectionValue(
				battles.CorrectionDamage,
				2048,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Water().ToString())),
		)
	}
	if weather == WeatherRainy {
		return battles.NewBattleCorrectionValues(
			battles.NewBattleCorrectionValue(
				battles.CorrectionDamage,
				6144,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Water().ToString())),
			battles.NewBattleCorrectionValue(
				battles.CorrectionDamage,
				2048,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Fire().ToString())),
		)
	}
	return nil
}

type WeatherType string

const (
	WeatherNormal    WeatherType = "Normal"
	WeatherSunny     WeatherType = "Sunny"
	WeatherRainy     WeatherType = "Rainy"
	WeatherHailstone WeatherType = "Hailstone"
	WeatherSandstorm WeatherType = "Sandstorm"
)
