package battles

import (
	"github.com/Symthy/PokeRest/pokeRest/common/fmath"
)

type BattleCorrectionValue struct {
	target           CorrectionTarget
	value            float32
	triggerCondition *TriggerCondition // Todo: 無しという状態を持つようようにした方が良い？
}

func NewBattleCorrectionValue(target CorrectionTarget, value float32, condition *TriggerCondition) *BattleCorrectionValue {
	if isInvalidValue(value) {
		return nil
	}
	return &BattleCorrectionValue{
		target:           CorrectionTarget(target),
		value:            value,
		triggerCondition: condition,
	}
}

func (c BattleCorrectionValue) Apply(
	input uint16, battleDataSet IPokemonBattleDataSet, side BattleSideType) uint16 {
	if c.triggerCondition == nil { // Todo
		return fmath.Round[uint16](float64(input) * float64(c.value) / 4096)
	}
	if c.triggerCondition.isSatisfy(battleDataSet, side) {
		return fmath.Round[uint16](float64(input) * float64(c.value) / 4096)
	}
	return input
}

func (c BattleCorrectionValue) equalTarget(target CorrectionTarget) bool {
	return c.target == target
}

func (c BattleCorrectionValue) AnyEqualTarget(targets ...CorrectionTarget) bool {
	for _, target := range targets {
		if c.target == target {
			return true
		}
	}
	return false
}

func isInvalidValue(rate float32) bool {
	return rate < 0
}

type CorrectionType string

const (
	CorrectionTypeAttack  CorrectionType = "Attack"
	CorrectionTypeDefense CorrectionType = "Defense"
)

func (c CorrectionType) String() string {
	return string(c)
}
