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
	file, err := os.Open("input.csv")
	require.NoError(t, err)
	defer file.Close()
	reader := bufio.NewReader(file)

	customers, err := parseFile(reader)
	require.NoError(t, err)
	assert.Equal(t, 1_000_000, len(customers))
}

func Benchmark_parseFile(b *testing.B) {
	benchmarkParseFile(b, parseFile)
}

func Benchmark_parseFileWorker(b *testing.B) {
	benchmarkParseFile(b, parseFileWorker)
}

func Benchmark_parseFileGoroutines(b *testing.B) {
	benchmarkParseFile(b, parseFileGoroutines)
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
