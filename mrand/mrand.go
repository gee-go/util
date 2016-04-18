package mrand

import (
	"math/rand"
	"time"
)

// Seed the default rand source.
func Seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func NewSource() rand.Source {
	return rand.NewSource(time.Now().UnixNano())
}

func New() *rand.Rand {
	return rand.New(NewSource())
}
