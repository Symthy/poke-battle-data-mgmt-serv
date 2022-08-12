package dto

import (
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
	"github.com/stretchr/testify/assert"
)

func TestToSchemaBattleEffects(t *testing.T) {
	cases := map[string]struct {
		input    *battles.BattleEffects
		expected *mixin.BattleEffects
	}{
		"empty": {
			input:    battles.NewEmptyBattleEffects(),
			expected: mixin.NewBattleEffects(),
		},
	}

	for testcase, tt := range cases {
		t.Run(testcase, func(t *testing.T) {
			assert.EqualValues(t, tt.expected, toSchemaBattleEffects(tt.input))
		})
	}
}
