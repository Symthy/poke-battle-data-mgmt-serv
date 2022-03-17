package enum

import "database/sql/driver"

type CorrectionTarget string

const (
	// 技威力
	PhysicalMove CorrectionTarget = "physicalMove"
	SpecialMove  CorrectionTarget = "specialMove"
	// 攻撃力
	Attack        CorrectionTarget = "attack"
	SpecialAttack CorrectionTarget = "specialAttack"
	// 攻撃威力
	AttackPower        CorrectionTarget = "attackPower"
	SpecialAttackPower CorrectionTarget = "specialAttackPower"
	// 防御力
	Deffence        CorrectionTarget = "deffence"
	SpecialDeffence CorrectionTarget = "specialDeffence"
	// ダメージ
	Damage CorrectionTarget = "damage"
	// 重さ
	Weight CorrectionTarget = "weight"
	// タイプ一致
	TypeMatch CorrectionTarget = "typeMatch"
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

type ApplyCorrectionType string

const (
	Multiply ApplyCorrectionType = "multiply"
	Replace  ApplyCorrectionType = "replace"
	fixed    ApplyCorrectionType = "fixed"
)

func (c *ApplyCorrectionType) Scan(value interface{}) error {
	*c = ApplyCorrectionType(value.([]byte))
	return nil
}

func (c ApplyCorrectionType) Value() (driver.Value, error) {
	return string(c), nil
}

func (c ApplyCorrectionType) String() string {
	return string(c)
}

type ConditionEntry string

const (
	ConditionPokemonType       ConditionEntry = "pokemonType"
	ConditionTypeCompatibility ConditionEntry = "typeCompatibility"
	ConditionIsTypeMatch       ConditionEntry = "isTypeMatch"
	ConditionGender            ConditionEntry = "gender"
)

func (c *ConditionEntry) Scan(value interface{}) error {
	*c = ConditionEntry(value.([]byte))
	return nil
}

func (c ConditionEntry) Value() (driver.Value, error) {
	return string(c), nil
}

func (c ConditionEntry) String() string {
	return string(c)
}
