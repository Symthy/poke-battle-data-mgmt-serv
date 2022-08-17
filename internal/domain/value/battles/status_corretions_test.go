package battles

import (
	"testing"

	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/stretchr/testify/assert"
)

var _ IPokemonBattleDataSet = (*mockDataSetForStatusCorrections)(nil)

type mockDataSetForStatusCorrections struct {
	EmptyPokemonBattleBaseParams
	EmptyTriggerConditionParams
}

func (_m mockDataSetForStatusCorrections) AttackPokemonActualValues() *value.PokemonActualValues {
	return value.NewPokemonActualValues(120, 120, 120, 120, 120, 110)
}

func TestStatusCorrectionsApply(t *testing.T) {
	cases := map[string]struct {
		inputCorrections *StatusCorrections
		expectedA        uint16
		expectedB        uint16
		expectedC        uint16
		expectedD        uint16
		expectedS        uint16
	}{
		"correction target is exist": {
			inputCorrections: NewStatusCorrections(
				NewBattleCorrectionValues(
					NewDefaultCorrectionValue(CorrectionStatusA, Correction6144, nil),
					NewDefaultCorrectionValue(CorrectionStatusB, Correction5734, nil),
					NewDefaultCorrectionValue(CorrectionStatusC, Correction5325, nil),
					NewDefaultCorrectionValue(CorrectionStatusD, Correction5448, nil),
					NewDefaultCorrectionValue(CorrectionStatusS, Correction2732, nil),
				),
				BattleAttackSide,
			),
			expectedA: 180,
			expectedB: 168,
			expectedC: 156,
			expectedD: 160, // 159.6 小数点以下五超え
			expectedS: 73,  // 73.36 小数点以下五未満
		},
	}

	mockDataSet := mockDataSetForStatusCorrections{}
	for testcaseName, tt := range cases {
		t.Run(testcaseName, func(t *testing.T) {
			assert.Equal(t, tt.expectedA, tt.inputCorrections.ApplyA(mockDataSet))
			assert.Equal(t, tt.expectedB, tt.inputCorrections.ApplyB(mockDataSet))
			assert.Equal(t, tt.expectedC, tt.inputCorrections.ApplyC(mockDataSet))
			assert.Equal(t, tt.expectedD, tt.inputCorrections.ApplyD(mockDataSet))
			assert.Equal(t, tt.expectedS, tt.inputCorrections.ApplyS(mockDataSet))
		})
	}
}
