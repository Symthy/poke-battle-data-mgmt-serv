package mixin

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
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

	tests := []struct {
		inputBytes []byte
		expected   mixin.BattleEffects
	}{
		{
			inputBytes: []byte(`{}`),
			expected:   mixin.BattleEffects{},
		},
		{
			inputBytes: []byte(`{
				"corrections": [],
				"overrides": [],
			}`),
			expected: mixin.BattleEffects{
				Corrections: []mixin.Correction{},
				Overrides:   []mixin.Override{},
			},
		},
		{
			inputBytes: []byte(`{
				"corrections": [
					{
						"target": "PhysicalMove",
						"value": 1.2,
						"triggerCondition": {
							"entry": "PokemonType",
							"value": "Normal"
						}
					},
					{
						"target": "SpecialMove",
						"value": 1.2,
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
			expected: mixin.BattleEffects{
				Corrections: []mixin.Correction{
					{
						Target: enum.CorrectionPhysicalMove,
						Value:  1.2,
						TriggerCondition: &mixin.TriggerCondition{
							Entry: enum.ConditionPokemonType,
							Value: enum.Normal.String(),
						},
					},
				},
				Overrides: []mixin.Override{
					{
						Target: enum.OverridePokemonType,
						Value:  enum.Flying.String(),
						TriggerCondition: &mixin.TriggerCondition{
							Entry: enum.ConditionPokemonType,
							Value: enum.Normal.String(),
						},
					},
				},
			},
		},
		{
			inputBytes: []byte(`{
				"corrections": [
					{
						"target": "Damage",
						"value": 1.5,
					},
				],
				"overrides": [
					{
						"target": "FixedDamage",
						"value": "50",
					}
				]}`),
			expected: mixin.BattleEffects{
				Corrections: []mixin.Correction{
					{
						Target: enum.CorrectionDamage,
						Value:  1.5,
					},
				},
				Overrides: []mixin.Override{
					{
						Target: enum.OverrideFixedDamage,
						Value:  "50",
					},
				},
			},
		},
	}

	for i, tt := range tests {
		actual := mixin.BattleEffects{}
		err := json.Unmarshal(tt.inputBytes, &actual)
		if err != nil {
			fmt.Printf("error: [index:%d] %#v\n", i, err)
			suite.Fail("failure json unmarshal")
			return
		}
		assert.EqualValues(suite.T(), tt.expected, actual)
	}
}
