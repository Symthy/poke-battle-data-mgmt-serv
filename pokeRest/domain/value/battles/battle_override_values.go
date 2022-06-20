package battles

import "github.com/Symthy/PokeRest/pokeRest/common/lists"

type BattleOverrideValues struct {
	values  []*BattleOverrideValue
	targets []OverrideTarget
}

func NewBattleOverrideValues(values ...*BattleOverrideValue) *BattleOverrideValues {
	if len(values) == 0 {
		return &BattleOverrideValues{
			values:  []*BattleOverrideValue{},
			targets: []OverrideTarget{},
		}
	}
	return &BattleOverrideValues{
		values: values,
		targets: lists.Map(values, func(v *BattleOverrideValue) OverrideTarget {
			return v.target
		}),
	}
}

func (o *BattleOverrideValues) Merge(overrides *BattleOverrideValues) {
	o.values = append(o.values, overrides.values...)
	o.targets = append(o.targets, overrides.targets...)
}
