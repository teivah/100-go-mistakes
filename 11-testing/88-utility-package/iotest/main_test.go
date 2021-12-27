package main

import (
	"strings"
	"testing"
	"testing/iotest"
)

func TestLowerCaseReader(t *testing.T) {
	err := iotest.TestReader(
		&LowerCaseReader{reader: strings.NewReader("aBcDeFgHiJ")},
		[]byte("acegi"),
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFoo(t *testing.T) {
	err := foo(iotest.TimeoutReader(
		strings.NewReader(randomString(1024)),
	))
	if err != nil {
		t.Fatal(err)
	}
}

func randomString(i int) string {
	return string(make([]byte, i))
}
