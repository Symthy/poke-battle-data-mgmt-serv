package battles

import (
	"github.com/Symthy/PokeRest/pokeRest/common/fmath"
)

type BattleCorrectionValue struct {
	target           CorrectionTarget
	value            float32
	triggerCondition *TriggerCondition // Todo: 無しという状態を持つようようにした方が良い？
}

func NewBattleCorrectionValue(target string, value float32, condition *TriggerCondition) *BattleCorrectionValue {
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
	return rate <= 0
}

type CorrectionTarget string

const (
	// 技威力への補正
	PhysicalMoveCorrection CorrectionTarget = "PhysicalMove"
	SpecialMoveCorrection  CorrectionTarget = "SpecialMove"
	// 威力補正
	PhysicalPowerCorrection CorrectionTarget = "AttackPower"
	SpecialPowerCorrection  CorrectionTarget = "SpecialAttackPower"
	// ステータス補正
	AttackCorrection         CorrectionTarget = "Attack"
	SpecialAttackCorrection  CorrectionTarget = "SpecialAttack"
	DefenseCorrection        CorrectionTarget = "Defense"
	SpecialDefenseCorrection CorrectionTarget = "SpecialDefense"
	SpeedCorrection          CorrectionTarget = "Speed"
	WeightCorrection         CorrectionTarget = "Weight"
	// ダメージ
	DamageCorrection CorrectionTarget = "Damage"
	// 補正なし
	NoneCorrection CorrectionTarget = ""
)

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
		AttackCorrection,
		SpecialAttackCorrection,
		DefenseCorrection,
		SpecialDefenseCorrection,
		SpeedCorrection,
		WeightCorrection,
	}
}

func GetPowerCorrectionTargets() []CorrectionTarget {
	return []CorrectionTarget{
		PhysicalPowerCorrection,
		SpecialPowerCorrection,
	}
}

func GetMovePowerCorrectionTargets() []CorrectionTarget {
	return []CorrectionTarget{
		PhysicalMoveCorrection,
		SpecialMoveCorrection,
	}
}

func GetDamageCorrectionTargets() []CorrectionTarget {
	return []CorrectionTarget{
		DamageCorrection,
	}
}
