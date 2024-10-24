package xsync

import (
	"arena"
	"sync"
)

type BytesPool struct {
	cap  int
	pool *sync.Pool
	area *arena.Arena
}

func NewBytesHeapSyncPool(cap int) *BytesPool {
	return &BytesPool{
		cap:  cap,
		area: nil,
		pool: &sync.Pool{
			New: func() any {
				pool := make([]byte, 0, cap)
				return &pool
			},
		},
	}
}

func NewBytesArenaSyncPool(cap int) *BytesPool {
	area := arena.NewArena()

	return &BytesPool{
		cap:  cap,
		area: area,
		pool: &sync.Pool{
			New: func() any {
				pool := arena.MakeSlice[byte](area, 0, cap)
				return &pool
			},
		},
	}
}

func (p *BytesPool) Get() *[]byte {
	return p.pool.Get().(*[]byte)
}

func (p *BytesPool) Put(pool *[]byte) {
	// check if the pool is the correct size
	if cap(*pool) != p.cap {
		return
	}
	// reset the pool to be empty
	*pool = (*pool)[:0]
	// put the pool back in the pool
	p.pool.Put(pool)
}
