package battles

import (
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/stretchr/testify/assert"
)

type mockTriggerConditionParamsForCorrectionValue struct {
	EmptyTriggerConditionParams
}

func (m mockTriggerConditionParamsForCorrectionValue) AttackPokemonTypeOfFirst() value.PokemonType {
	return value.Normal()
}

func TestApply(t *testing.T) {
	cases := map[string]struct {
		crrectionValue *BattleCorrectionValue
		input          uint16
		expected       uint16
	}{
		"non condition (always applied)": {
			crrectionValue: NewDefaultCorrectionValue(CorrectionPhysicalMove, 2048, nil),
			input:          200,
			expected:       100,
		},
		"satisfy condition": {
			crrectionValue: NewDefaultCorrectionValue(CorrectionPhysicalMove, 2048,
				NewTriggerCondition(ConditionPokemonType, value.Normal().ToString())),
			input:    200,
			expected: 100,
		},
		"non satisfy condition": {
			crrectionValue: NewDefaultCorrectionValue(CorrectionPhysicalMove, 2048,
				NewTriggerCondition(ConditionPokemonType, value.Rock().ToString())),
			input:    200,
			expected: 200,
		},
	}

	for testcaseName, tt := range cases {
		t.Run(testcaseName, func(t *testing.T) {
			actual := tt.crrectionValue.Apply(200, mockTriggerConditionParamsForCorrectionValue{}, BattleAttackSide)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
