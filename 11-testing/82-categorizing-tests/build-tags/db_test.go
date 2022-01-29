//go:build integration
// +build integration

package db

import (
	"os"
	"testing"
)

func TestInsert1(t *testing.T) {
	// ...
}

func TestInsert2(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip("skipping integration test")
	}

	// ...
}
