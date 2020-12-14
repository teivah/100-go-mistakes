package concurrency

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const size = 10_000_000

func Test_mergeSort(t *testing.T) {
	var (
		input    = []int{5, 8, 9, 5, 0, 10, 1, 6}
		expected = []int{0, 1, 5, 5, 6, 8, 9, 10}
	)

	type args struct {
		f     func(s []int)
		input []int
	}
	tests := map[string]struct {
		args     args
		expected []int
	}{
		"sequential": {
			args: args{
				f:     mergesortSequential,
				input: input,
			},
			expected: expected,
		},
		"concurrent v1": {
			args: args{
				f:     mergeSortConcurrentV1,
				input: input,
			},
			expected: expected,
		},
		"concurrent v2": {
			args: args{
				f:     mergeSortConcurrentV2,
				input: input,
			},
			expected: expected,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			tt.args.f(tt.args.input)
			assert.Equal(t, tt.expected, tt.args.input)
		})
	}
}

func Benchmark_mergeSortSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := random(size)
		b.StartTimer()
		mergesortSequential(s)
		b.StopTimer()
	}
}

func Benchmark_mergeSortConcurrentV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := random(size)
		b.StartTimer()
		mergeSortConcurrentV1(s)
		b.StopTimer()
	}
}

func Benchmark_mergeSortConcurrentV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := random(size)
		b.StartTimer()
		mergeSortConcurrentV2(s)
		b.StopTimer()
	}
}

func random(n int) []int {
	s := make([]int, n)

	src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(src)

	for i := 0; i < n; i++ {
		s[i] = rand.Intn(n)
	}

	return s
}
