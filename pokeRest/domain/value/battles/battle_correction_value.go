package battles

import (
	"github.com/Symthy/PokeRest/pokeRest/common/fmath"
)

type BattleCorrectionValue struct {
	target                 CorrectionTarget
	value                  float32
	triggerCondition       *TriggerCondition // Todo: 無しという状態を持つようようにした方が良い？
	decimalPointCalculator func(float64) uint16
}

func NewDefaultCorrectionValue(target CorrectionTarget, value float32, condition *TriggerCondition) *BattleCorrectionValue {
	if isInvalidValue(value) {
		return nil
	}
	return &BattleCorrectionValue{
		target:                 CorrectionTarget(target),
		value:                  value,
		triggerCondition:       condition,
		decimalPointCalculator: fmath.Round[uint16],
	}
}

func NewCorrectionValue(target CorrectionTarget, value float32, condition *TriggerCondition, decimalPointCalculator func(float64) uint16) *BattleCorrectionValue {
	if isInvalidValue(value) {
		return nil
	}
	return &BattleCorrectionValue{
		target:                 CorrectionTarget(target),
		value:                  value,
		triggerCondition:       condition,
		decimalPointCalculator: decimalPointCalculator,
	}
}

func NewBattleNonCorrectionValue() *BattleCorrectionValue {
	return &BattleCorrectionValue{
		target: CorrectionNone,
	}
}

func isInvalidValue(value float32) bool {
	return value < 0
}

func (c BattleCorrectionValue) Apply(
	input uint16, battleDataSet TriggerConditionParams, side BattleSideType) uint16 {
	if c.triggerCondition == nil { // Todo
		return c.decimalPointCalculator(float64(input) * float64(c.value) / 4096.0)
	}
	if c.triggerCondition.isSatisfy(battleDataSet, side) {
		return c.decimalPointCalculator(float64(input) * float64(c.value) / 4096.0)
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

type CorrectionType string

const (
	CorrectionTypeAttack  CorrectionType = "Attack"
	CorrectionTypeDefense CorrectionType = "Defense"
)

func (c CorrectionType) String() string {
	return string(c)
}
