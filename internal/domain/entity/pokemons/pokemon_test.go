package pokemons

import (
	"testing"

	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/stretchr/testify/assert"
)

func TestCalculateActualValueH(t *testing.T) {
	cases := []struct {
		baseStats       int
		individualValue int
		effortValue     int
		expected        int
	}{
		{
			baseStats:       80,
			individualValue: 31,
			effortValue:     0,
			expected:        155,
		},
		{
			baseStats:       80,
			individualValue: 0,
			effortValue:     0,
			expected:        140,
		},
		{
			baseStats:       80,
			individualValue: 31,
			effortValue:     4,
			expected:        156,
		},
		{
			baseStats:       80,
			individualValue: 31,
			effortValue:     252,
			expected:        187,
		},
	}

	for _, tt := range cases {
		effort, _ := value.NewEffortValue(uint64(tt.effortValue), value.H)
		actual := int(calculateActualValueH(
			value.NewPokemonBaseStats(uint64(tt.baseStats)),
			uint8(tt.individualValue),
			*effort,
		))

		assert.Equal(t, tt.expected, actual)
	}
}

func TestCalculateActualValueABCDS(t *testing.T) {
	cases := []struct {
		baseStats       int
		individualValue int
		effortValue     int
		expected        int
	}{
		{
			baseStats:       100,
			individualValue: 31,
			effortValue:     0,
			expected:        120,
		},
		{
			baseStats:       100,
			individualValue: 0,
			effortValue:     0,
			expected:        105,
		},
		{
			baseStats:       100,
			individualValue: 31,
			effortValue:     4,
			expected:        121,
		},
		{
			baseStats:       100,
			individualValue: 31,
			effortValue:     252,
			expected:        152,
		},
	}

	for _, tt := range cases {
		effort, _ := value.NewEffortValue(uint64(tt.effortValue), value.A)
		actual := int(calculateActualValueABCDS(
			value.NewPokemonBaseStats(uint64(tt.baseStats)),
			uint8(tt.individualValue),
			*effort,
		))

		assert.Equal(t, tt.expected, actual)
	}
}
