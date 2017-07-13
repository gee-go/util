package vec

import (
	"testing"

	"github.com/DataDog/testify/assert"
)

func TestRoundEven(t *testing.T) {
	tests := []struct {
		v   float64
		exp float64
	}{
		{-.1, 0},
		{-.6, -1},
		{.1, 0},
		{.2, 0},
		{.6, 1},

		// .5
		{-.5, 0},
		{.5, 0},
		{1.5, 2},
		{2.5, 2},
		{3.5, 4},
	}

	for _, tc := range tests {
		act := roundEven(tc.v)
		assert.Equal(t, tc.exp, act, "%+v", tc)
	}
}
