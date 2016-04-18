package mrand

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAlphaString(t *testing.T) {
	a := require.New(t)

	out := AlphaBytes(NewSource(), 40)
	a.Len(out, 40)
}

func BenchmarkAlphaString(b *testing.B) {
	// a := require.New(t)
	src := NewSource()
	var s []byte
	for i := 0; i < b.N; i++ {
		s = AlphaBytes(src, 40)
	}

	if len(s) != 40 {
		b.Fail()
	}

}
