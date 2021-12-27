package main

import "testing"

var global string

func BenchmarkConcatV1(b *testing.B) {
	var local string
	s := getInput()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = concat1(s)
	}
	global = local
}

func BenchmarkConcatV2(b *testing.B) {
	var local string
	s := getInput()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = concat2(s)
	}
	global = local
}

func BenchmarkConcatV3(b *testing.B) {
	var local string
	s := getInput()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = concat3(s)
	}
	global = local
}

func getInput() []string {
	n := 1_000
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = string(make([]byte, 1_000))
	}
	return s
}
