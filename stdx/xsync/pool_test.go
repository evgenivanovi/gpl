package xsync_test

import (
	"testing"

	"github.com/evgenivanovi/gpl/stdx/xsync"
)

func BenchmarkBytesHeapSyncPool(b *testing.B) {
	pool := xsync.NewBytesHeapSyncPool(1024)

	b.Run("GetPut", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			buf := pool.Get()
			pool.Put(buf)
		}
	})

	b.Run("ParallelGetPut", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				buf := pool.Get()
				pool.Put(buf)
			}
		})
	})
}

func BenchmarkBytesArenaSyncPool(b *testing.B) {
	pool := xsync.NewBytesArenaSyncPool(1024)

	b.Run("GetPut", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			buf := pool.Get()
			pool.Put(buf)
		}
	})

	b.Run("ParallelGetPut", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				buf := pool.Get()
				pool.Put(buf)
			}
		})
	})
}
