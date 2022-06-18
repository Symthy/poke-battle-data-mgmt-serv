package battles

type BattleEffects struct {
	corrections *BattleCorrectionValues
	overrides   *BattleOverrideValues
}

func NewEmptyBattleEffects() *BattleEffects {
	return &BattleEffects{
		corrections: NewBattleCorrectionValues(),
		overrides:   NewBattleOverrideValues(),
	}
}

func NewBattleEffects(corrections *BattleCorrectionValues, overrides *BattleOverrideValues) *BattleEffects {
	return &BattleEffects{corrections, overrides}
}

func (b *BattleEffects) Merge(effects *BattleEffects) {
	b.corrections.Merge(effects.corrections)
	b.overrides.Merge(effects.overrides)
}

func (b BattleEffects) Corrections() *BattleCorrectionValues {
	return b.corrections
}

func (b BattleEffects) Overrides() *BattleOverrideValues {
	return b.overrides
}

func (b BattleEffects) ApplyCorrection(
	value uint16, target CorrectionTarget, data IPokemonBattleDataSet, side BattleSideType) uint16 {
	return b.corrections.Apply(value, target, data, side)
}
