package kv

import "sync"

type shard struct {
	mu    sync.Mutex
	data  map[KEY_TYPE]VAL_TYPE
	stats Stats
}

type Sharded struct {
	shards    []shard
	numShards int
}

func (m *Sharded) hashKey(k KEY_TYPE) int {
	return fnv64a(uint64(k), m.numShards)
}

func (m *Sharded) PerShardStats() []Stats {
	sp := make([]Stats, len(m.shards))

	for i := range m.shards {
		m.shards[i].mu.Lock()
		sp[i] = m.shards[i].stats
		sp[i].size = len(m.shards[i].data)
		m.shards[i].mu.Unlock()
	}

	return sp
}

func (m *Sharded) Stats() Stats {
	var total Stats

	for i := range m.shards {
		m.shards[i].mu.Lock()
		ss := m.shards[i].stats
		total.misses += ss.misses
		total.hits += ss.hits
		total.size += len(m.shards[i].data)
		m.shards[i].mu.Unlock()
	}

	return total
}

func (m *Sharded) Set(k KEY_TYPE, v VAL_TYPE) {
	s := m.hashKey(k)

	m.shards[s].mu.Lock()
	m.shards[s].data[k] = v
	m.shards[s].mu.Unlock()
}

func (m *Sharded) Get(k KEY_TYPE) (VAL_TYPE, bool) {
	s := m.hashKey(k)

	m.shards[s].mu.Lock()
	v, found := m.shards[s].data[k]
	if found {
		m.shards[s].stats.hits++
	} else {
		m.shards[s].stats.misses++
	}
	m.shards[s].mu.Unlock()
	return v, found
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
