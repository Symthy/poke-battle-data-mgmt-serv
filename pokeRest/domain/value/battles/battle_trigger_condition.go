package battles

import (
	"strconv"

	"github.com/Symthy/PokeRest/pokeRest/domain/entity/types"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type ConditionStisfiesResolver func(TriggerConditionParams, BattleSideType) bool

type TriggerCondition struct {
	entry    ConditionEntry
	value    string
	resolver ConditionStisfiesResolver
}

func NewTriggerCondition(entry ConditionEntry, conditionValue string) *TriggerCondition {
	return &TriggerCondition{
		entry:    entry,
		value:    conditionValue,
		resolver: buildConditionStisfiesResolver(entry, conditionValue),
	}
}

func (t TriggerCondition) isSatisfy(data TriggerConditionParams, side BattleSideType) bool {
	return t.resolver(data, side)
}

func (t TriggerCondition) Notify(note ITriggerConditionNotification) {
	note.SetEntry(t.entry)
	note.SetConditionValue(t.value)
}

func buildConditionStisfiesResolver(entry ConditionEntry, conditionValue string) ConditionStisfiesResolver {
	// Todo
	switch entry {
	case ConditionPokemonType:
		return buildPokemonTypeSatisfiesResolver(conditionValue)
	case ConditionTypeCompatibility:
		return buildTypeCompatibilitySatisfiesResolver(conditionValue)
	case ConditionTypeMatch:
		return buildTypeMatchSatisfiesResolver(conditionValue)
	default:
		return func(arg TriggerConditionParams, side BattleSideType) bool {
			return false
		}
	}
}

func buildPokemonTypeSatisfiesResolver(targetType string) ConditionStisfiesResolver {
	targetPokemonType := value.NewPokemonType(targetType)
	return func(data TriggerConditionParams, side BattleSideType) bool {
		switch side {
		case BattleAttackSide:
			return data.AttackPokemonTypeOfFirst().Equals(targetPokemonType) ||
				data.AttackPokemonTypeOfSecond().Equals(targetPokemonType)
		case BattleDefenseSide:
			return data.DefensePokemonTypeOfFirst().Equals(targetPokemonType) ||
				data.DefensePokemonTypeOfSecond().Equals(targetPokemonType)
		default:
			return false
		}
	}
}

func buildTypeCompatibilitySatisfiesResolver(conditionDamageRate string) ConditionStisfiesResolver {
	return func(data TriggerConditionParams, side BattleSideType) bool {
		conditionDamageRateInt, err := strconv.Atoi(conditionDamageRate)
		if err != nil {
			return false
		}
		switch side {
		case BattleDefenseSide:
			compatibilityTable := types.NewTypeCompatibilityTable()
			damageRate := compatibilityTable.ResolveDamageRate(data.MovePokemonType(), data.DefensePokemonTypeOfFirst(),
				data.DefensePokemonTypeOfSecond())
			return damageRate >= float32(conditionDamageRateInt)
		default:
			return false
		}
	}
}

func buildTypeMatchSatisfiesResolver(targetType string) ConditionStisfiesResolver {
	targetPokemonType := value.NewPokemonType(targetType)
	return func(data TriggerConditionParams, side BattleSideType) bool {
		switch side {
		case BattleAttackSide:
			return data.AttackPokemonTypeOfFirst().Equals(targetPokemonType) ||
				data.AttackPokemonTypeOfSecond().Equals(targetPokemonType)
		case BattleDefenseSide:
			return data.DefensePokemonTypeOfFirst().Equals(targetPokemonType) ||
				data.DefensePokemonTypeOfSecond().Equals(targetPokemonType)
		default:
			return false
		}
	}
}
