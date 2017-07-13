package vec

import "math"

// roundEven is a bad implementation of the default
// rounding method used in numpy.
// XXX: not well tested
func roundEven(v float64) int {
	i, f := math.Modf(v)

	m := math.Copysign(0.5, v)

	// 0.5 - goes to nearest even.
	// tolerance was arbitrarily picked.
	if int(i)%2 == 0 && .0001 > math.Abs(m-f) {
		return int(i)
	}

	return int(v + (m * 1.001))
}
