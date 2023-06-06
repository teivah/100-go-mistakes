package main

import "testing"

var global int64

const n = 1_000_000

func BenchmarkLinkedList(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nodes := make([]node, n)
		for i := 0; i < n-1; i++ {
			nodes[i].next = &nodes[i+1]
		}
		b.StartTimer()
		local = linkedList(&nodes[0])
	}
	global = local
}

func BenchmarkSum2(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := make([]int64, n)
		b.StartTimer()
		local = sum2(s)
	}
	global = local
}
