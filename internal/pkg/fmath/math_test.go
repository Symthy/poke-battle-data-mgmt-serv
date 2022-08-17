package fmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundUpIfDecimalGreaterFive(t *testing.T) {
	cases := map[string]struct {
		input    float64
		expected uint32
	}{
		"decimal is greater than 5": {
			input:    1.6,
			expected: 2,
		},
		"decimal is 5 or less": {
			input:    1.5,
			expected: 1,
		},
	}

	for testcase, tt := range cases {
		t.Run(testcase, func(t *testing.T) {
			actual := RoundUpIfDecimalGreaterFive[uint32](tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
