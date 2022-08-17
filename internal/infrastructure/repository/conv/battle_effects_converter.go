package conv

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/mixin"
	"github.com/Symthy/PokeRest/internal/domain/value/battles"
)

func toBattleEffects(effects *mixin.BattleEffects) *battles.BattleEffects {
	corrections := []*battles.BattleCorrectionValue{}
	for _, correction := range effects.Corrections {
		triggerCondition := convertTriggerCondition(correction.TriggerCondition)
		correctionValue := battles.NewDefaultCorrectionValue(
			battles.NewCorrectionTarget(correction.Target.String()),
			correction.Value,
			triggerCondition,
		)
		if correctionValue != nil {
			corrections = append(corrections, correctionValue)
		}
	}

	overrides := []*battles.BattleOverrideValue{}
	for _, override := range effects.Overrides {
		triggerCondition := convertTriggerCondition(override.TriggerCondition)
		overrideValue := battles.NewBattleOverrideValue(
			override.Target.String(),
			override.Value,
			triggerCondition,
		)
		if overrideValue != nil {
			overrides = append(overrides, overrideValue)
		}
	}

	result := battles.NewBattleEffects(battles.NewBattleCorrectionValues(corrections...), battles.NewBattleOverrideValues(overrides...))
	return result
}

func convertTriggerCondition(triggerCondition *mixin.TriggerCondition) *battles.TriggerCondition {
	if triggerCondition == nil {
		return nil
	}
	result := battles.NewTriggerCondition(
		battles.NewConditionEntry(triggerCondition.Entry.String()),
		triggerCondition.Value,
	)
	return result
}
