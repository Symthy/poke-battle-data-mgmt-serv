package enum

import "database/sql/driver"

type CorrectionTarget string

const (
	// 技威力
	CorrectionPhysicalMove CorrectionTarget = "PhysicalMove"
	CorrectionSpecialMove  CorrectionTarget = "SpecialMove"
	// 攻撃力
	CorrectionAttack        CorrectionTarget = "Attack"
	CorrectionSpecialAttack CorrectionTarget = "SpecialAttack"
	// 攻撃威力
	AttackPower                  CorrectionTarget = "AttackPower"
	CorrectionSpecialAttackPower CorrectionTarget = "SpecialAttackPower"
	// 防御力
	CorrectionDeffence        CorrectionTarget = "Deffence"
	CorrectionSpecialDeffence CorrectionTarget = "SpecialDeffence"
	// ダメージ
	CorrectionDamage CorrectionTarget = "Damage"
	// 重さ
	CorrectionWeight CorrectionTarget = "Weight"

	// 補正なし
	NoneCorrection CorrectionTarget = ""
)

func (c *CorrectionTarget) Scan(value interface{}) error {
	*c = CorrectionTarget(value.([]byte))
	return nil
}

func (c CorrectionTarget) Value() (driver.Value, error) {
	return string(c), nil
}

func (c CorrectionTarget) String() string {
	return string(c)
}
