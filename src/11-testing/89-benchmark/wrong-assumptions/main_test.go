package main

import (
	"sync/atomic"
	"testing"
)

func BenchmarkAtomicStoreInt32(b *testing.B) {
	var v int32
	for i := 0; i < b.N; i++ {
		atomic.StoreInt32(&v, 1)
	}
}

func BenchmarkAtomicStoreInt64(b *testing.B) {
	var v int64
	for i := 0; i < b.N; i++ {
		atomic.StoreInt64(&v, 1)
	}
}
