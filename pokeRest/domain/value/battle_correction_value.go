package value

import "github.com/Symthy/PokeRest/pokeRest/common/fmath"

type BattleCorrectionValue struct {
	target           CorrectionTarget
	value            int
	triggerCondition *TriggerCondition // Todo: 無しという状態を持つようようにした方が良い？
}

func NewBattleCorrectionValue(target string, value int, condition *TriggerCondition) *BattleCorrectionValue {
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
	input int, battleDataSet IPokemonBattleDataSet, side BattleSideType) int {
	if c.triggerCondition == nil { // Todo
		return fmath.Round(float64(input*c.value) / 4096)
	}
	if c.triggerCondition.isSatisfy(battleDataSet, side) {
		return fmath.Round(float64(input*c.value) / 4096)
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

func isInvalidValue(rate int) bool {
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
	DefenceCorrection        CorrectionTarget = "Defence"
	SpecialDefenceCorrection CorrectionTarget = "SpecialDefence"
	SpeedCorrection          CorrectionTarget = "Speed"
	WeightCorrection         CorrectionTarget = "Weight"
	// ダメージ
	DamageCorrection CorrectionTarget = "Damage"
	// 補正なし
	NoneCorrection CorrectionTarget = ""
)

type CorrectionType string

const (
	CorrectionTypeAttack  CorrectionType = "Attack"
	CorrectionTypeDefence CorrectionType = "Defence"
)

func GetStatusCorrectionTargets() []CorrectionTarget {
	return []CorrectionTarget{
		AttackCorrection,
		SpecialAttackCorrection,
		DefenceCorrection,
		SpecialDefenceCorrection,
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
