package enum

import "database/sql/driver"

type OverrideTarget string

const (
	// ダメージ
	OverrideFixedDamage OverrideTarget = "FixedDamage"
	// タイプ一致補正
	OverrideTypeMatchCorrection OverrideTarget = "TypeMatch"
	// タイプ
	OverridePokemonType OverrideTarget = "PokemonType"
)

func (c *OverrideTarget) Scan(value interface{}) error {
	*c = OverrideTarget(value.([]byte))
	return nil
}

func (c OverrideTarget) Value() (driver.Value, error) {
	return string(c), nil
}

func (c OverrideTarget) String() string {
	return string(c)
}
