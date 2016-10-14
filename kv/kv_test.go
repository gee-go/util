package kv

import (
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/lazybeaver/xorshift"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	m := NewSharded(32)

	m.Set(1, "abc")
	assert := assert.New(t)
	v, found := m.Get(1)
	assert.True(found)
	assert.Equal("abc", v)

	var wg sync.WaitGroup

	for w := 0; w < 40; w++ {
		wg.Add(1)
		go func() {
			seed := uint64(time.Now().UnixNano())
			rnd := xorshift.NewXorShift128Plus(seed)
			for i := 0; i < 1<<16; i++ {
				m.Set(KEY_TYPE(rnd.Next()), "a")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	avg := 0
	for _, s := range m.PerShardStats() {
		avg += s.size

		// fmt.Printf("shard %d len=%d\n", i, )
	}

	avg = avg / len(m.shards)
	for _, s := range m.PerShardStats() {
		assert.InDelta(avg, s.size, 10*(float64(avg)/100))
	}
	// pp.Println(m.PerShardStats())
	// fmt.Println(avg)

}

func BenchmarkMapParallel(b *testing.B) {
	cases := map[string]Map{
		"simple": NewSimple(),
	}
	seed := uint64(time.Now().UnixNano())
	for _, sn := range []int{1, 64, 256} {
		cases[strconv.Itoa(sn)+"-sharded"] = NewSharded(sn)
	}

	for name, m := range cases {
		b.Run(name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				var i KEY_TYPE
				rnd := xorshift.NewXorShift128Plus(seed)
				for pb.Next() {
					m.Set(KEY_TYPE(rnd.Next()), "a")
					i++
				}
			})
		})
	}

	for name, m := range cases {
		b.Run(name+"get", func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				var i KEY_TYPE
				var val VAL_TYPE
				// rnd := xorshift.NewXorShift128Plus(seed)
				// var found bool
				for pb.Next() {
					v, found := m.Get(i % 10000)
					if found {
						val = v
					}
					i++
				}

				if val != "a" && val != "" {
					b.Fail()
				}
			})
		})
	}
}
