package kv

import "sync"

type shard struct {
	mu   sync.Mutex
	data map[KEY_TYPE]VAL_TYPE
}

type Sharded struct {
	shards    []shard
	numShards int
}

func (m *Sharded) hashKey(k KEY_TYPE) int {
	return jumpHash(uint64(k), m.numShards)
}

func (m *Sharded) Set(k KEY_TYPE, v VAL_TYPE) {
	s := m.hashKey(k)

	m.shards[s].mu.Lock()
	m.shards[s].data[k] = v
	m.shards[s].mu.Unlock()
}

func (m *Sharded) Get(k KEY_TYPE) VAL_TYPE {
	s := m.hashKey(k)

	m.shards[s].mu.Lock()
	v := m.shards[s].data[k]
	m.shards[s].mu.Unlock()
	return v
}

func NewSharded(numShards int) *Sharded {
	m := &Sharded{
		shards:    make([]shard, numShards),
		numShards: numShards,
	}
	for i := range m.shards {
		m.shards[i].data = make(map[KEY_TYPE]VAL_TYPE)
	}

	return m
}
