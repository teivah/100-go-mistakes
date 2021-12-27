package main

import "testing"

const n = 1_000_000

var global map[int]struct{}

func BenchmarkMapWithoutSize(b *testing.B) {
	var local map[int]struct{}
	for i := 0; i < b.N; i++ {
		m := make(map[int]struct{})
		for j := 0; j < n; j++ {
			m[j] = struct{}{}
		}
	}
	global = local
}

func BenchmarkMapWithSize(b *testing.B) {
	var local map[int]struct{}
	for i := 0; i < b.N; i++ {
		m := make(map[int]struct{}, n)
		for j := 0; j < n; j++ {
			m[j] = struct{}{}
		}
	}
	global = local
}
