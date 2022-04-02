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
				return data.AttackPokemonTypeOfFirst().NameEN() == value || data.AttackPokemonTypeOfSecond().NameEN() == value
			case BattleDefenseSide:
				return data.DefensePokemonTypeOfFirst().NameEN() == value || data.DefensePokemonTypeOfSecond().NameEN() == value
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

func (c ConditionEntry) String() string {
	return string(c)
}

type IPokemonBattleDataSet interface {
	AttackPokemonTypeOfFirst() PokemonType
	AttackPokemonTypeOfSecond() PokemonType
	AttackPokemonActualValueS() int
	DefensePokemonTypeOfFirst() PokemonType
	DefensePokemonTypeOfSecond() PokemonType
	DefensePokemonActualValueS() int
	MovePokemonType() PokemonType
	HasItemAttackSide() bool
	HasItemDefenseSide() bool
}
