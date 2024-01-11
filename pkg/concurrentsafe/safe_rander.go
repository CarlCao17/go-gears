package concurrentsafe

import (
	"math/rand"
	"sync"
	"sync/atomic"
)

var defaultSize = 128

type SafeRander struct {
	pos   uint32
	rands []rand.Rand
	locks []sync.Mutex
}

func (sr *SafeRander) Intn(n int) int {
	x := atomic.AddUint32(&sr.pos, 1)
	x %= uint32(len(sr.rands))
	sr.locks[x].Lock()
	defer sr.locks[x].Unlock()
	return sr.rands[x].Intn(n)
}

func NewSafeRander(size int) *SafeRander {
	if size <= 0 {
		size = defaultSize
	}
	return &SafeRander{
		rands: make([]rand.Rand, size),
		locks: make([]sync.Mutex, size),
	}
}
