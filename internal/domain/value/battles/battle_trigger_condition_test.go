package battles

import (
	"testing"

	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/stretchr/testify/assert"
)

type mockTriggerConditionParamsForAttackType struct {
	EmptyTriggerConditionParams
}

func (m mockTriggerConditionParamsForAttackType) AttackPokemonTypeOfFirst() value.PokemonType {
	return value.Steel()
}

type mockTriggerConditionParamsForDefenseType struct {
	EmptyTriggerConditionParams
}

func (m mockTriggerConditionParamsForDefenseType) DefensePokemonTypeOfFirst() value.PokemonType {
	return value.Steel()
}

func TestBuildPokemonTypeSatisfiesResolver(t *testing.T) {
	cases := map[string]struct {
		triggerCondition *TriggerCondition
		conditionParams  TriggerConditionParams
		side             BattleSideType
		expected         bool
	}{
		"pokemon type attack side match": {
			triggerCondition: NewTriggerCondition(ConditionPokemonType, value.Steel().ToString()),
			conditionParams:  mockTriggerConditionParamsForAttackType{},
			side:             BattleAttackSide,
			expected:         true,
		},
		"pokemon type attack side non match": {
			triggerCondition: NewTriggerCondition(ConditionPokemonType, value.Normal().ToString()),
			conditionParams:  mockTriggerConditionParamsForAttackType{},
			side:             BattleAttackSide,
			expected:         false,
		},
		"pokemon type defense side match": {
			triggerCondition: NewTriggerCondition(ConditionPokemonType, value.Steel().ToString()),
			conditionParams:  mockTriggerConditionParamsForDefenseType{},
			side:             BattleDefenseSide,
			expected:         true,
		},
		"pokemon type defense side non match": {
			triggerCondition: NewTriggerCondition(ConditionPokemonType, value.Normal().ToString()),
			conditionParams:  mockTriggerConditionParamsForDefenseType{},
			side:             BattleDefenseSide,
			expected:         false,
		},
	}

	for testCaseName, tt := range cases {
		t.Run(testCaseName, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.triggerCondition.resolver(tt.conditionParams, tt.side))
		})
	}
}
