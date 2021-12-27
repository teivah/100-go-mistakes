package main

import (
	"math/rand"
	"testing"
	"time"
)

var global []int

func Benchmark_sequentialMergesort(b *testing.B) {
	var local []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := getRandomElements()
		b.StartTimer()

		sequentialMergesort(input)
		local = input
	}
	global = local
}

func Benchmark_parallelMergesortV1(b *testing.B) {
	var local []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := getRandomElements()
		b.StartTimer()

		parallelMergesortV1(input)
		local = input
	}
	global = local
}

func Benchmark_parallelMergesortV2(b *testing.B) {
	var local []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := getRandomElements()
		b.StartTimer()

		parallelMergesortV2(input)
		local = input
	}
	global = local
}

func getRandomElements() []int {
	n := 10_000
	res := make([]int, n)
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	for i := 0; i < n; i++ {
		res[i] = rnd.Int()
	}
	return res
}
