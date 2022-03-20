package conv

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

func toBattleEffects(effects *mixin.BattleEffects) value.BattleEffects {
	corrections := []value.BattleCorrectionValue{}
	for _, correction := range effects.Corrections {
		triggerCondition := convertTriggerCondition(correction.TriggerCondition)
		correctionValue := value.NewBattleCorrectionValue(
			correction.Target.String(),
			correction.Value,
			triggerCondition,
		)
		if correctionValue != nil {
			corrections = append(corrections, *correctionValue)
		}
	}

	overrides := []value.BattleOverrideValue{}
	for _, override := range effects.Overrides {
		triggerCondition := convertTriggerCondition(override.TriggerCondition)
		overrideValue := value.NewBattleOverrideValue(
			override.Target.String(),
			override.Value,
			triggerCondition,
		)
		if overrideValue != nil {
			overrides = append(overrides, *overrideValue)
		}
	}

	result := value.NewBattleEffects(value.NewBattleCorrectionValues(corrections), value.NewBattleOverrideValues([]value.BattleOverrideValue{}))
	return result
}

func convertTriggerCondition(triggerCondition *mixin.TriggerCondition) *value.TriggerCondition {
	if triggerCondition == nil {
		return nil
	}
	result := value.NewTriggerCondition(
		triggerCondition.Entry.String(),
		triggerCondition.Sign.String(),
		triggerCondition.Value,
	)
	return result
}
