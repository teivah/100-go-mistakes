package main

import "testing"

var global int64

const n = 1_000_000

func BenchmarkSumFoo(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := make([]Foo, n)
		b.StartTimer()
		local = sumFoo(s)
	}
	global = local
}

func BenchmarkSumBar(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bar := Bar{
			a: make([]int64, n),
			b: make([]int64, n),
		}
		b.StartTimer()
		local = sumBar(bar)
	}
	global = local
}
