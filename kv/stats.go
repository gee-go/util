package kv

type Stats struct {
	misses uint64
	hits   uint64
	size   int
}
