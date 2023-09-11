package main

import "testing"

const n = 1_000_000

var globalResult1 Result1

func BenchmarkCount1(b *testing.B) {
	var local Result1
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		inputs := make([]Input, n)
		b.StartTimer()
		local = count1(inputs)
	}
	globalResult1 = local
}

var globalResult2 Result2

func BenchmarkCount2(b *testing.B) {
	var local Result2
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		inputs := make([]Input, n)
		b.StartTimer()
		local = count2(inputs)
	}
	globalResult2 = local
}
