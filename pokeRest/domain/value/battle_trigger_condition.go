package value

type ApplicableResolver func(IPokemonBattleDataSet, BattleSideType) bool

type TriggerCondition struct {
	resolver ApplicableResolver
}

func NewTriggerCondition(entry string, conditionValue string) *TriggerCondition {
	ent := ConditionEntry(entry)
	return &TriggerCondition{
		resolver: getApplicableResolver(ent, conditionValue),
	}
}

func getApplicableResolver(entry ConditionEntry, value string) ApplicableResolver {
	// Todo
	if entry == ConditionPokemonType {
		return func(data IPokemonBattleDataSet, side BattleSideType) bool {
			switch side {
			case BattleAttackSide:
				return data.AttackPokemonTypeOfFirst() == value || data.AttackPokemonTypeOfSecond() == value
			case BattleDefenceSide:
				return data.DefencePokemonTypeOfFirst() == value || data.DefencePokemonTypeOfSecond() == value
			default:
				return false
			}
		}
	}
	return func(arg IPokemonBattleDataSet, side BattleSideType) bool {
		return false
	}
}

func (t TriggerCondition) isSatisfy(data IPokemonBattleDataSet, side BattleSideType) bool {
	return t.resolver(data, side)
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
	DefencePokemonTypeOfFirst() string
	DefencePokemonTypeOfSecond() string
	DefencePokemonActualValueS() string
	MovePokemonType() string
	HasItemAttackSide() bool
	HasItemDefenceSide() bool
}
