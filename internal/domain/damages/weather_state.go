package damages

import (
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/battles"
)

type WeatherType string

const (
	WeatherNormal    WeatherType = "Normal"
	WeatherSunny     WeatherType = "Sunny"
	WeatherRainy     WeatherType = "Rainy"
	WeatherHailstone WeatherType = "Hailstone"
	WeatherSandstorm WeatherType = "Sandstorm"
)

type WeatherState struct {
	weather     WeatherType
	corrections *battles.BattleCorrectionValues
}

func NewNormalWeatherState() *WeatherState {
	return &WeatherState{
		weather:     WeatherNormal,
		corrections: resolveWeatherCorrection(WeatherNormal),
	}
}

func NewWeatherState(weather string) *WeatherState {
	return &WeatherState{
		weather:     WeatherType(weather),
		corrections: resolveWeatherCorrection(WeatherType(weather)),
	}
}

func (w WeatherState) correctedValue(p battles.IPokemonBattleDataSet) uint16 {
	if w.weather == WeatherNormal {
		return battles.Correction4096
	}
	return w.corrections.Apply(battles.Correction4096, battles.CorrectionDamage, nil, battles.BattleAttackSide)
}

func resolveWeatherCorrection(weather WeatherType) *battles.BattleCorrectionValues {
	if weather == WeatherSunny {
		return battles.NewBattleCorrectionValues(
			battles.NewDefaultCorrectionValue(
				battles.CorrectionDamage,
				battles.Correction6144,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Fire().ToString())),
			battles.NewDefaultCorrectionValue(
				battles.CorrectionDamage,
				battles.Correction2048,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Water().ToString())),
		)
	}
	if weather == WeatherRainy {
		return battles.NewBattleCorrectionValues(
			battles.NewDefaultCorrectionValue(
				battles.CorrectionDamage,
				battles.Correction6144,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Water().ToString())),
			battles.NewDefaultCorrectionValue(
				battles.CorrectionDamage,
				battles.Correction2048,
				battles.NewTriggerCondition(battles.ConditionPokemonType, value.Fire().ToString())),
		)
	}
	return battles.NewBattleCorrectionValues()
}
