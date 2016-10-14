package kv

import "testing"

var hash_val int

func BenchmarkHash(b *testing.B) {
	b.Run("jump", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hash_val = jumpHash(uint64(i)%50000, 256)
		}
	})

	b.Run("xorshiftMult64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hash_val = xorshiftMult64(uint64(i), 256)
		}
	})

	b.Run("fnv64a", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hash_val = fnv64a(uint64(i), 256)
		}
	})

	b.Run("javaSmear", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hash_val = javaSmear(uint64(i), 256)
		}
	})
}
