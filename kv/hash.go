package kv

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
