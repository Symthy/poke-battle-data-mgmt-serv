package value

type TriggerCondition struct {
	entry ConditionEntry
	sign  ConditionSign
	value string
}

func NewTriggerCondition(entry string, sign string, value string) *TriggerCondition {
	return &TriggerCondition{
		entry: ConditionEntry(entry),
		sign:  ConditionSign(sign),
		value: value,
	}
}

type ConditionEntry string

const (
	ConditionPokemonType       ConditionEntry = "pokemonType"
	ConditionTypeCompatibility ConditionEntry = "typeCompatibility"
	ConditionIsTypeMatch       ConditionEntry = "isTypeMatch"
	ConditionGender            ConditionEntry = "gender"
)

type ConditionSign string

const (
	StringMatchSign ConditionSign = "match"
	EqualSign       ConditionSign = "=="
	GraterThanSign  ConditionSign = ">"
)
