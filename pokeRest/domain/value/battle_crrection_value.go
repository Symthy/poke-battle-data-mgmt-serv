package value

type BattleCorrectionValue struct {
	target           CorrectionTarget
	value            float32
	triggerCondition *TriggerCondition
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

func (c BattleCorrectionValue) ApplyCorrection(inputTarget CorrectionTarget, input float32) float32 {
	if c.target != inputTarget {
		return input
	}
	return input * c.value
}

func isInvalidValue(rate float32) bool {
	return rate == 1 || rate <= 0
}

type CorrectionTarget string

const (
	// 技威力
	PhysicalMoveCorrection CorrectionTarget = "PhysicalMove"
	SpecialMoveCorrection  CorrectionTarget = "SpecialMove"
	// 攻撃力
	AttackCorrection        CorrectionTarget = "Attack"
	SpecialAttackCorrection CorrectionTarget = "SpecialAttack"
	// 攻撃威力
	AttackPowerCorrection        CorrectionTarget = "AttackPower"
	SpecialAttackPowerCorrection CorrectionTarget = "SpecialAttackPower"
	// 防御力
	DeffenceCorrection        CorrectionTarget = "Deffence"
	SpecialDeffenceCorrection CorrectionTarget = "SpecialDeffence"
	// ダメージ
	DamageCorrection CorrectionTarget = "Damage"
	// 重さ
	WeightCorrection CorrectionTarget = "Weight"
	// 補正なし
	NoneCorrection CorrectionTarget = ""
)
