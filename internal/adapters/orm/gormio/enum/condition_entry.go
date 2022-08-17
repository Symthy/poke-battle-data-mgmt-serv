package enum

import "database/sql/driver"

type ConditionEntry string

const (
	ConditionPokemonType       ConditionEntry = "PokemonType"
	ConditionTypeCompatibility ConditionEntry = "TypeCompatibility"
	ConditionIsTypeMatch       ConditionEntry = "IsTypeMatch"
	ConditionGender            ConditionEntry = "Gender"
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

type ConditionSign string

const (
	StringMatchSign ConditionSign = "match"
	EqualSign       ConditionSign = "=="
	GraterThanSign  ConditionSign = ">"
)

func (c *ConditionSign) Scan(value interface{}) error {
	*c = ConditionSign(value.([]byte))
	return nil
}

func (c ConditionSign) Value() (driver.Value, error) {
	return string(c), nil
}

func (c ConditionSign) String() string {
	return string(c)
}
