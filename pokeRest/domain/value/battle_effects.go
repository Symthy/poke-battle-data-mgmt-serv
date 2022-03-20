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
