package value

type BattleEffects struct {
	corrections BattleCorrectionValues
	overrides   BattleOverrideValues
}

func NewBattleEffects(corrections BattleCorrectionValues, overrides BattleOverrideValues) BattleEffects {
	return BattleEffects{corrections, overrides}
}

func (b *BattleEffects) Merge(effects BattleEffects) {
	b.corrections.Merge(effects.corrections)
	b.overrides.Merge(effects.overrides)
}

func (b BattleEffects) Corrections() BattleCorrectionValues {
	return b.corrections
}

func (b BattleEffects) Overrides() BattleOverrideValues {
	return b.overrides
}

func (b BattleEffects) ApplyCorrection(value float32, target CorrectionTarget) float32 {
	return b.corrections.Apply(value, target)
}

type BattleSideType string

const (
	BattleAttackSide  BattleSideType = "AttackSide"
	BattleDefenceSide BattleSideType = "DefenceSide"
)
