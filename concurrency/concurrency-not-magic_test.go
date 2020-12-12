package concurrency

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_parseFile(t *testing.T) {
	type args struct {
		f func(reader *bufio.Reader) ([]customer, error)
	}
	tests := map[string]struct {
		args args
	}{
		"sequential": {
			args: args{
				f: parseFileSequential,
			},
		},
		"concurrent v1": {
			args: args{
				f: parseFileConcurrentV1,
			},
		},
		"concurrent v2": {
			args: args{
				f: parseFileConcurrentV2,
			},
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			file, err := os.Open("input.csv")
			require.NoError(t, err)
			defer file.Close()
			reader := bufio.NewReader(file)

			customers, err := tt.args.f(reader)
			require.NoError(t, err)
			assert.Equal(t, 1_000_000, len(customers))
		})
	}
}

//
//func Benchmark_parseFileSequential(b *testing.B) {
//	benchmarkParseFile(b, parseFileSequential)
//}
//
//func Benchmark_parseFileConcurrentV1(b *testing.B) {
//	benchmarkParseFile(b, parseFileConcurrentV1)
//}

func Benchmark_parseFileConcurrentV2(b *testing.B) {
	benchmarkParseFile(b, parseFileConcurrentV2)
}

func benchmarkParseFile(b *testing.B, f func(reader *bufio.Reader) ([]customer, error)) {
	file, err := os.Open("input.csv")
	require.NoError(b, err)
	csv, _ := ioutil.ReadAll(file)
	reader := bufio.NewReader(bytes.NewReader(csv))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(reader)
	}
}
