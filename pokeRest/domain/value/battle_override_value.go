package value

type BattleOverrideValues struct {
	values  []BattleOverrideValue
	targets []OverrideTarget
}

func NewBattleOverrideValues(values []BattleOverrideValue) BattleOverrideValues {
	return BattleOverrideValues{
		values: values,
	}
}

func (o *BattleOverrideValues) Merge(overrides BattleOverrideValues) {
	o.values = append(o.values, overrides.values...)
	o.targets = append(o.targets, overrides.targets...)
}

type BattleOverrideValue struct {
	target           OverrideTarget
	value            string
	triggerCondition *TriggerCondition
}

func NewBattleOverrideValue(target string, value string, condition *TriggerCondition) *BattleOverrideValue {
	if value == "" {
		return nil
	}
	return &BattleOverrideValue{
		target:           OverrideTarget(target),
		value:            value,
		triggerCondition: condition,
	}
}

type OverrideTarget string

const (
	// ダメージ
	OverrideFixedDamage OverrideTarget = "FixedDamage"
	// タイプ一致補正
	OverrideTypeMatchCorrection OverrideTarget = "TypeMatch"
	// タイプ
	OverridePokemonType OverrideTarget = "PokemonType"
)
