package vec

import (
	"fmt"
	"math"
	"sort"
)

// Interpolation method to use when the desired quantile lies between two data points i < j.
// See https://docs.scipy.org/doc/numpy-1.12.0/reference/generated/numpy.percentile.html
type Interpolation uint8

const (

	// Linear: i + (j - i) * fraction, where fraction is the fractional part of the index surrounded by i and j
	Linear Interpolation = iota

	// Lower: i
	Lower

	// Higher: j
	Higher

	// Nearest: i or j, whichever is nearest.
	Nearest

	// Midpoint: (i + j) / 2
	Midpoint
)

// Quantile calculates the quantile like numpy does (except q is from [0,1] not [0, 100])
func Quantile(data []float64, q float64, interp Interpolation) float64 {
	if q < 0 || q > 1 {
		return math.NaN()
	}

	if !sort.Float64sAreSorted(data) {
		sort.Float64s(data)
	}

	n := float64(len(data) - 1)
	idx := q * n

	switch interp {
	case Linear:
		fidx, cw := math.Modf(idx)
		cidx, fw := math.Min(fidx+1, n), 1-cw

		cv, fv := data[int(cidx)], data[int(fidx)]

		return cv*cw + fv*fw
	case Lower:
		return data[int(idx)]
	case Higher:
		return data[int(math.Ceil(idx))]
	case Nearest:
		return data[int(roundEven(idx))]
	case Midpoint:
		c := data[int(math.Ceil(idx))]
		f := data[int(math.Floor(idx))]

		return (c + f) / 2
	}

	panic(fmt.Errorf("unknown interp %v", interp))
}
