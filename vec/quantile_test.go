package vec

import (
	"testing"

	"github.com/DataDog/testify/assert"
)

func TestQuantile(t *testing.T) {
	arange := func(n int) []float64 {
		out := make([]float64, n)
		for i := range out {
			out[i] = float64(i)
		}
		return out
	}

	r10 := arange(10)
	r11 := arange(11)

	tests := []struct {
		data   []float64
		q      float64
		method Interpolation
		exp    float64
	}{
		{r10, .5, Linear, 4.5},
		{r10, .5, Lower, 4},
		{r10, .5, Higher, 5},
		{r10, 1, Higher, 9},
		{r11, 1, Higher, 10},
		{r10, .51, Midpoint, 4.5},
		{r11, .51, Midpoint, 5.5},
		{r11, .50, Midpoint, 5},
		{r10, .51, Nearest, 5},
		{r10, .49, Nearest, 4},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.exp, Quantile(tc.data, tc.q, tc.method), "%+v", tc)
	}

}
