package main

import (
	"strings"
	"testing"
)

func TestCountEmptyLines(t *testing.T) {
	emptyLines, err := countEmptyLines(strings.NewReader(
		`foo
			bar

			baz
			`))
	// Test logic
	_ = emptyLines
	_ = err
}
