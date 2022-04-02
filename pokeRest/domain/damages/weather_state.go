package damages

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type WeatherState struct {
	weather     WeatherType
	corrections *value.BattleCorrectionValues
}

func NewWeatherState(weather string) WeatherState {
	return WeatherState{
		weather:     WeatherType(weather),
		corrections: resolveWeatherCorrection(WeatherType(weather)),
	}
}

func (w WeatherState) ApplyCorrection(damage int, data IPokemonBattleDataSet) int {
	if w.weather == WeatherNormal {
		return damage
	}
	return w.corrections.Apply(damage, value.DamageCorrection, data, value.BattleAttackSide)
}

func resolveWeatherCorrection(weather WeatherType) *value.BattleCorrectionValues {
	if weather == WeatherSunny {
		return value.NewBattleCorrectionValues(
			value.NewBattleCorrectionValue(
				value.DamageCorrection.String(),
				6144,
				value.NewTriggerCondition(value.ConditionPokemonType.String(), value.Fire().NameEN())),
			value.NewBattleCorrectionValue(
				value.DamageCorrection.String(),
				2048,
				value.NewTriggerCondition(value.ConditionPokemonType.String(), value.Water().NameEN())),
		)
	}
	if weather == WeatherRainy {
		return value.NewBattleCorrectionValues(
			value.NewBattleCorrectionValue(
				value.DamageCorrection.String(),
				6144,
				value.NewTriggerCondition(value.ConditionPokemonType.String(), value.Water().NameEN())),
			value.NewBattleCorrectionValue(
				value.DamageCorrection.String(),
				2048,
				value.NewTriggerCondition(value.ConditionPokemonType.String(), value.Fire().NameEN())),
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
