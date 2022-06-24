package battles

type BattleOverrideValue struct {
	target           OverrideTarget
	value            string
	triggerCondition *TriggerCondition
}

func NewBattleOverrideValue(target string, value string, condition *TriggerCondition) *BattleOverrideValue {
	return &BattleOverrideValue{
		target:           OverrideTarget(target),
		value:            value,
		triggerCondition: condition,
	}
}

func (v BattleOverrideValue) IsMatch(target OverrideTarget) bool {
	return v.target == target
}

func (v BattleOverrideValue) Apply() string {
	return v.value
}
