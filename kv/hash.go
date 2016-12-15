package kv

const (
	offset32 uint32 = 2166136261
	offset64 uint64 = 14695981039346656037
	prime32  uint32 = 16777619
	prime64  uint64 = 1099511628211
)

type hashFn func(uint64, int) int

const numShards = 256

// see http://arxiv.org/pdf/1406.2294.pdf
func jumpHash(k uint64, numBuckets int) int {

	var b int64 = -1
	var j int64

	for j < int64(numBuckets) {
		b = j
		k = k*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((k>>33)+1)))
	}

	return int(b)
}

func javaSmear(hashCode uint64, numBuckets int) int {
	// used by guava slotted hash map
	hashCode = hashCode ^ (hashCode >> 32)
	hashCode ^= (hashCode >> 20) ^ (hashCode >> 12)
	hashCode = hashCode ^ (hashCode >> 7) ^ (hashCode >> 4)

	return int(hashCode & (uint64(numBuckets) - 1))
}

func xorshiftMult64(x uint64, numBuckets int) int {
	x ^= x >> 12 // a
	x ^= x << 25 // b
	x ^= x >> 27 // c
	return int((x * 2685821657736338717) & (uint64(numBuckets) - 1))
}

func fnv64a(v uint64, numBuckets int) int {
	h := offset32

	h ^= uint32(byte(v))
	h *= prime32
	h ^= uint32(byte(v >> 8))
	h *= prime32
	h ^= uint32(byte(v >> 16))
	h *= prime32
	h ^= uint32(byte(v >> 24))
	h *= prime32
	h ^= uint32(byte(v >> 32))
	h *= prime32
	h ^= uint32(byte(v >> 40))
	h *= prime32
	h ^= uint32(byte(v >> 48))
	h *= prime32
	h ^= uint32(byte(v >> 56))
	h *= prime32

	return int(h) & (numBuckets - 1)
}

func fnv64a2(v uint64, numBuckets int) int {
	h := offset64

	for i := uint64(0); i <= 56; i += 8 {
		h ^= (v >> i) & 8
		h *= prime64
	}
	// h = (prime32 * h) ^ uint32(byte(v))
	// h = (prime32 * h) ^ uint32(byte(v>>8))
	// h = (prime32 * h) ^ uint32(byte(v>>16))
	// h = (prime32 * h) ^ uint32(byte(v>>24))
	// h = (prime32 * h) ^ uint32(byte(v>>32))
	// h = (prime32 * h) ^ uint32(byte(v>>40))
	// h = (prime32 * h) ^ uint32(byte(v>>48))
	// h = (prime32 * h) ^ uint32(byte(v>>56))

	// TODO - premature optimization - xorshift to 8 bytes
	return int(h) & (numBuckets - 1)
}
