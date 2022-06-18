package battles

import (
	"github.com/Symthy/PokeRest/pokeRest/common/fmath"
	"github.com/Symthy/PokeRest/pokeRest/common/lists"
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

type CorrectionTarget string

const (
	// 技威力への補正
	CorrectionPhysicalMove CorrectionTarget = "PhysicalMove" // 物理技
	CorrectionSpecialMove  CorrectionTarget = "SpecialMove"
	// 威力補正
	CorrectionPhysicalPower CorrectionTarget = "PhysicalPower"
	CorrectionSpecialPower  CorrectionTarget = "SpecialPower"
	// ステータス補正
	CorrectionStatusA CorrectionTarget = "StatusA"
	CorrectionStatusC CorrectionTarget = "StatusC"
	CorrectionStatusB CorrectionTarget = "StatusB"
	CorrectionStatusD CorrectionTarget = "StatusD"
	CorrectionStatusS CorrectionTarget = "StatusS"
	// ダメージ
	CorrectionDamage CorrectionTarget = "Damage"
	// その他
	CorrectionWeight CorrectionTarget = "Weight"
	CorrectionNone   CorrectionTarget = ""
)

func allCorrectionTarget() []CorrectionTarget {
	return []CorrectionTarget{
		CorrectionPhysicalPower,
		CorrectionSpecialMove,
		CorrectionPhysicalPower,
		CorrectionSpecialPower,
		CorrectionStatusA,
		CorrectionStatusC,
		CorrectionStatusB,
		CorrectionStatusD,
		CorrectionStatusS,
		CorrectionWeight,
	}
}

func NewCorrectionTarget(target string) CorrectionTarget {
	if lists.Contains(allCorrectionTarget(), target) {
		return CorrectionTarget(target)
	}
	return CorrectionNone
}

func (c CorrectionTarget) String() string {
	return string(c)
}

type CorrectionType string

const (
	CorrectionTypeAttack  CorrectionType = "Attack"
	CorrectionTypeDefense CorrectionType = "Defense"
)

func (c CorrectionType) String() string {
	return string(c)
}

func GetStatusCorrectionTargets() []CorrectionTarget {
	return []CorrectionTarget{
		CorrectionStatusA,
		CorrectionStatusC,
		CorrectionStatusB,
		CorrectionStatusD,
		CorrectionStatusS,
		CorrectionWeight,
	}
}

func GetPowerCorrectionTargets() []CorrectionTarget {
	return []CorrectionTarget{
		CorrectionPhysicalPower,
		CorrectionSpecialPower,
	}
}

func GetMovePowerCorrectionTargets() []CorrectionTarget {
	return []CorrectionTarget{
		CorrectionPhysicalMove,
		CorrectionSpecialMove,
	}
}

func GetDamageCorrectionTargets() []CorrectionTarget {
	return []CorrectionTarget{
		CorrectionDamage,
	}
}
