package battles

import "github.com/Symthy/PokeRest/internal/pkg/fmath"

type CorrectionNum = uint16

const (
	Correction8192 CorrectionNum = 8192 // 2.00
	Correction7372 CorrectionNum = 7372 // 1.80
	Correction6553 CorrectionNum = 6553 // 1.60
	Correction6144 CorrectionNum = 6144 // 1.50
	Correction5734 CorrectionNum = 5734 // 1.40
	Correction5448 CorrectionNum = 5448 // 1.33
	Correction5325 CorrectionNum = 5325 // 1.30
	Correction5120 CorrectionNum = 5120 // 1.25
	Correction4915 CorrectionNum = 4915 // 1.20
	Correction4505 CorrectionNum = 4505 // 1.10
	Correction4096 CorrectionNum = 4096 // 1.00
	Correction3072 CorrectionNum = 3072 // 0.75
	Correction2732 CorrectionNum = 2732 // 0.66
	Correction2048 CorrectionNum = 2048 // 0.50
	Correction1352 CorrectionNum = 1352 // 0.33
	Correction1024 CorrectionNum = 1024 // 0.25
)

type DecimalCalcType string

func (d DecimalCalcType) ToString() string {
	return string(d)
}

const (
	RoundUp                     DecimalCalcType = "up"
	RoundDown                   DecimalCalcType = "down"
	RoundUpIfDecimalGreaterFour DecimalCalcType = "upIfDecimalGreaterFour"
	RoundUpIfDecimalGreaterFive DecimalCalcType = "upIfDecimalGreaterFive"
)

type BattleCorrectionValue struct {
	target           CorrectionTarget
	value            CorrectionNum
	triggerCondition *TriggerCondition // Todo: 無しという状態を持つようようにした方が良い？
	decimalCalcType  DecimalCalcType   // 関数を持たせるとテストで等価比較できないため
}

func NewDefaultCorrectionValue(target CorrectionTarget, value CorrectionNum, condition *TriggerCondition) *BattleCorrectionValue {
	if isInvalidValue(value) {
		return NewNonCorrectionValue()
	}
	correctionTarget := CorrectionTarget(target)
	return &BattleCorrectionValue{
		target:           correctionTarget,
		value:            value,
		triggerCondition: condition,
		decimalCalcType:  overrideDecimalCalcType(correctionTarget, RoundUpIfDecimalGreaterFour),
	}
}

func NewCorrectionValue(
	target CorrectionTarget, value CorrectionNum, condition *TriggerCondition, decimalCalcType DecimalCalcType,
) *BattleCorrectionValue {
	if isInvalidValue(value) {
		return nil
	}
	correctionTarget := CorrectionTarget(target)
	return &BattleCorrectionValue{
		target:           correctionTarget,
		value:            value,
		triggerCondition: condition,
		decimalCalcType:  overrideDecimalCalcType(correctionTarget, RoundUpIfDecimalGreaterFour),
	}
}

func NewNonCorrectionValue() *BattleCorrectionValue {
	return &BattleCorrectionValue{
		target: CorrectionNone,
	}
}

func isInvalidValue(value CorrectionNum) bool {
	return value < 0
}

func overrideDecimalCalcType(target CorrectionTarget, decimalCalcType DecimalCalcType) DecimalCalcType {
	switch target {
	case CorrectionStatusA, CorrectionStatusB, CorrectionStatusC, CorrectionStatusD, CorrectionStatusS:
		return RoundUpIfDecimalGreaterFive
	default:
		return decimalCalcType
	}
}

func (c BattleCorrectionValue) Apply(
	input uint16, params TriggerConditionParams, side BattleSideType) uint16 {
	if c.triggerCondition == nil || c.triggerCondition.isSatisfy(params, side) {
		switch c.decimalCalcType {
		case RoundUp:
			return fmath.RoundUp[uint16](float64(float64(input) * float64(c.value) / float64(Correction4096)))
		case RoundDown:
			return fmath.RoundDown[uint16](float64(input) * float64(c.value) / float64(Correction4096))
		case RoundUpIfDecimalGreaterFour:
			return fmath.Round[uint16](float64(input) * float64(c.value) / float64(Correction4096))
		case RoundUpIfDecimalGreaterFive:
			return fmath.RoundUpIfDecimalGreaterFive[uint16](float64(input) * float64(c.value) / float64(Correction4096))
		default:
			return fmath.Round[uint16](float64(input) * float64(c.value) / float64(Correction4096))
		}
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

func (c BattleCorrectionValue) Notify(note ICorrectionValueNotification) {
	note.SetTarget(c.target)
	note.SetCorrectionValue(c.value)
	c.triggerCondition.Notify(note)
	note.SetDecimalCalcType(c.decimalCalcType)
}

type CorrectionType string

const (
	CorrectionTypeAttack  CorrectionType = "Attack"
	CorrectionTypeDefense CorrectionType = "Defense"
)

func (c CorrectionType) String() string {
	return string(c)
}
