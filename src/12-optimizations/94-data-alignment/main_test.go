package main

import "testing"

const n = 1_000_000

var global int64

func BenchmarkSum1(b *testing.B) {
	var local int64
	s := make([]Foo1, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = sum1(s)
	}
	global = local
}

func BenchmarkSum2(b *testing.B) {
	var local int64
	s := make([]Foo2, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = sum2(s)
	}
	global = local
}
