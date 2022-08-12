package battles

import (
	"github.com/Symthy/PokeRest/pokeRest/common/lists"
)

type correctionApplier func(value uint16) uint16

type correctionsApplier[T any] func(values T) T

func identify(value uint16) uint16 {
	return value
}

type BattleCorrectionValues struct {
	values  []*BattleCorrectionValue
	targets []CorrectionTarget
}

func NewBattleCorrectionValues(values ...*BattleCorrectionValue) *BattleCorrectionValues {
	if len(values) == 0 {
		return &BattleCorrectionValues{
			values:  []*BattleCorrectionValue{},
			targets: []CorrectionTarget{},
		}
	}
	return &BattleCorrectionValues{
		values: values,
		targets: lists.Map(values, func(v *BattleCorrectionValue) CorrectionTarget {
			return v.target
		}),
	}
}

func (b BattleCorrectionValues) get(targets ...CorrectionTarget) *BattleCorrectionValues {
	values := lists.Filter(b.values, func(correctionValue *BattleCorrectionValue) bool {
		return correctionValue.AnyEqualTarget(targets...)
	})
	return NewBattleCorrectionValues(values...)
}

func (c BattleCorrectionValues) IsEmpty() bool {
	return len(c.targets) == 0
}

func (c BattleCorrectionValues) Apply(
	input CorrectionNum, target CorrectionTarget, triggerConditions TriggerConditionParams, side BattleSideType) uint16 {
	result := input
	for _, correctionValue := range c.values {
		if target != correctionValue.target {
			continue
		}
		result = correctionValue.Apply(result, triggerConditions, side)
	}
	return result
}

func (c *BattleCorrectionValues) Merge(corrections *BattleCorrectionValues) {
	c.values = append(c.values, corrections.values...)
	c.targets = append(c.targets, corrections.targets...)
}

func (c BattleCorrectionValues) ToSlice() []*BattleCorrectionValue {
	return c.values
}
