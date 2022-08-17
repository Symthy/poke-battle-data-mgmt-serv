package battles

import "github.com/Symthy/PokeRest/internal/pkg/lists"

type BattleEffectsBuilder struct {
	correctionValues []*BattleCorrectionValue
	overrideValues   []*BattleOverrideValue
}

func NewBattleEffectsBuilder() *BattleEffectsBuilder {
	return &BattleEffectsBuilder{
		correctionValues: []*BattleCorrectionValue{},
		overrideValues:   []*BattleOverrideValue{},
	}
}

func (b *BattleEffectsBuilder) AddCorrectionValues(correctionValue ...*BattleCorrectionValue) {
	b.correctionValues = lists.Add(b.correctionValues, correctionValue...)
}

func (b *BattleEffectsBuilder) AddOverrideValues(overrideValues ...*BattleOverrideValue) {
	b.overrideValues = lists.Add(b.overrideValues, overrideValues...)
}

func (b BattleEffectsBuilder) Build() *BattleEffects {
	return NewBattleEffects(
		NewBattleCorrectionValues(b.correctionValues...),
		NewBattleOverrideValues(b.overrideValues...),
	)
}
