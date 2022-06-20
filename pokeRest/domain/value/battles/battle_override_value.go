package battles

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
