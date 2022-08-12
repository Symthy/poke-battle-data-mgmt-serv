package mixin

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BattleEffectsTestSuite struct {
	suite.Suite
}

// Before
func (suite *BattleEffectsTestSuite) SetupTest() {
}

// After
func (suite *BattleEffectsTestSuite) TearDownTest() {
}

func TestBattleEffectsTestSuite(t *testing.T) {
	suite.Run(t, new(BattleEffectsTestSuite))
}

func (suite BattleEffectsTestSuite) TestUnmarshal() {

	tests := map[string]struct {
		inputBytes []byte
		expected   BattleEffects
	}{
		"empty object": {
			inputBytes: []byte(`{}`),
			expected:   BattleEffects{},
		},
		"elements empty": {
			inputBytes: []byte(`{
				"corrections": [],
				"overrides": []
			}`),
			expected: BattleEffects{
				Corrections: []*Correction{},
				Overrides:   []*Override{},
			},
		},
		"data full set": {
			inputBytes: []byte(`{
				"corrections": [
					{
						"target": "PhysicalMove",
						"value": 4915,
						"triggerCondition": {
							"entry": "PokemonType",
							"value": "Normal"
						}
					},
					{
						"target": "SpecialMove",
						"value": 5325,
						"triggerCondition": {
							"entry": "PokemonType",
							"value": "Normal"
						}
					}
				],
				"overrides": [
					{
						"target": "PokemonType",
						"value": "Flying",
						"triggerCondition": {
							"entry": "PokemonType",
							"value": "Normal"
						}
					}
				]}`),
			expected: BattleEffects{
				Corrections: []*Correction{
					{
						Target: enum.CorrectionPhysicalMove,
						Value:  4915,
						TriggerCondition: &TriggerCondition{
							Entry: enum.ConditionPokemonType,
							Value: enum.Normal.String(),
						},
					},
					{
						Target: enum.CorrectionSpecialMove,
						Value:  5325,
						TriggerCondition: &TriggerCondition{
							Entry: enum.ConditionPokemonType,
							Value: enum.Normal.String(),
						},
					},
				},
				Overrides: []*Override{
					{
						Target: enum.OverridePokemonType,
						Value:  enum.Flying.String(),
						TriggerCondition: &TriggerCondition{
							Entry: enum.ConditionPokemonType,
							Value: enum.Normal.String(),
						},
					},
				},
			},
		},
		"trigger condition is not exist": {
			inputBytes: []byte(`{
				"corrections": [
					{
						"target": "Damage",
						"value": 6144
					}
				],
				"overrides": [
					{
						"target": "FixedDamage",
						"value": "50"
					}
				]}`),
			expected: BattleEffects{
				Corrections: []*Correction{
					{
						Target: enum.CorrectionDamage,
						Value:  6144,
					},
				},
				Overrides: []*Override{
					{
						Target: enum.OverrideFixedDamage,
						Value:  "50",
					},
				},
			},
		},
	}

	for testcaseName, tt := range tests {
		suite.T().Run(testcaseName, func(t *testing.T) {
			actual := BattleEffects{}
			err := json.Unmarshal(tt.inputBytes, &actual)
			if err != nil {
				fmt.Printf("error: %s: %#v\n", testcaseName, err)
				suite.Fail("failure json unmarshal")
				return
			}
			assert.EqualValues(t, tt.expected, actual)
		})
	}
}
