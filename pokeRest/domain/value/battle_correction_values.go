package value

import "github.com/Symthy/PokeRest/pokeRest/common/lists"

type BattleCorrectionValues struct {
	values  []BattleCorrectionValue
	targets []CorrectionTarget
}

func NewBattleCorrectionValues(values []BattleCorrectionValue) BattleCorrectionValues {
	targets := []CorrectionTarget{}
	for _, correction := range values {
		targets = append(targets, correction.target)
	}
	return BattleCorrectionValues{values, targets}
}

func (c BattleCorrectionValues) ApplyCorrections(correctionTarget CorrectionTarget, value float32) float32 {
	if !lists.Contains(c.targets, correctionTarget) {
		return value
	}
	result := value
	for _, correction := range c.values {
		result = correction.ApplyCorrection(correctionTarget, result)
	}
	return result
}

func (c *BattleCorrectionValues) Merge(corrections BattleCorrectionValues) {
	c.values = append(c.values, corrections.values...)
	c.targets = append(c.targets, corrections.targets...)
}
