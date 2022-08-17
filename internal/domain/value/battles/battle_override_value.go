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

func (o BattleOverrideValue) IsMatch(target OverrideTarget) bool {
	return o.target == target
}

func (o BattleOverrideValue) Apply() string {
	return o.value
}

func (o BattleOverrideValue) Notify(note IOverrideValueNotification) {
	note.SetTarget(o.target)
	note.SetOverrideValue(o.value)
	o.triggerCondition.Notify(note)
}
