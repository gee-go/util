package kv

import "sync"

type Simple struct {
	mu   sync.Mutex
	data map[KEY_TYPE]VAL_TYPE
}

func NewSimple() *Simple {
	return &Simple{data: make(map[KEY_TYPE]VAL_TYPE)}
}

func (m *Simple) Set(k KEY_TYPE, v VAL_TYPE) {
	m.mu.Lock()
	m.data[k] = v
	m.mu.Unlock()
}

func (m *Simple) Get(k KEY_TYPE) VAL_TYPE {
	m.mu.Lock()
	v := m.data[k]
	m.mu.Unlock()
	return v
}
