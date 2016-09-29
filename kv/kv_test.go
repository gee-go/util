package kv

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/lazybeaver/xorshift"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	m := NewSharded(256)

	m.Set(1, "abc")
	assert := assert.New(t)
	assert.Equal("abc", m.Get(1))

	var wg sync.WaitGroup

	for w := 0; w < 10; w++ {
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
	for i, s := range m.shards {
		fmt.Printf("shard %d len=%d\n", i, len(s.data))
	}

}

func BenchmarkMapParallel(b *testing.B) {
	cases := map[string]Map{
		"simple": NewSimple(),
	}

	for i := uint(3); i <= 10; i++ {
		sn := 1 << i
		cases[strconv.Itoa(sn)+"-sharded"] = NewSharded(sn)
	}

	for name, m := range cases {
		b.Run(name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				var i KEY_TYPE
				for pb.Next() {
					m.Set(i%10000, "a")
					i++
				}
			})
		})
	}
}
