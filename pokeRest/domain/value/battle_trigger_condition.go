package value

type ApplicableResolver func(IPokemonBattleDataSet) bool

type TriggerCondition struct {
	entry    ConditionEntry
	value    string
	resolver ApplicableResolver
}

func NewTriggerCondition(entry string, value string) *TriggerCondition {
	ent := ConditionEntry(entry)
	return &TriggerCondition{
		entry:    ent,
		value:    value,
		resolver: getApplicableResolver(ent, value),
	}
}

func getApplicableResolver(entry ConditionEntry, conditionValue string) ApplicableResolver {
	// Todo
	if entry == ConditionPokemonType {
		return func(arg IPokemonBattleDataSet) bool {
			return arg.AttackPokemonTypeOfFirst() == conditionValue
		}
	}
	return func(arg IPokemonBattleDataSet) bool {
		return false
	}
}

type ConditionEntry string

const (
	ConditionPokemonType       ConditionEntry = "pokemonType"
	ConditionTypeCompatibility ConditionEntry = "typeCompatibility"
	ConditionIsTypeMatch       ConditionEntry = "isTypeMatch"
	ConditionGender            ConditionEntry = "gender"
)

type IPokemonBattleDataSet interface {
	AttackPokemonTypeOfFirst() string
	AttackPokemonTypeOfSecond() string
	AttackPokemonActualValueS() string
	DefencePokemonActualValueS() string
	MovePokemonType() string
	HasItemAttackSide() bool
	HasItemDefenceSide() bool
}
