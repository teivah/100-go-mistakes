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

func TestFoo1(t *testing.T) {
	err := foo1(iotest.TimeoutReader(
		strings.NewReader(randomString(1024)),
	))
	if err != nil {
		t.Fatal(err)
	}
}

func TestFoo2(t *testing.T) {
	err := foo2(iotest.TimeoutReader(
		strings.NewReader(randomString(1024)),
	))
	if err != nil {
		t.Fatal(err)
	}
}

func randomString(i int) string {
	return string(make([]byte, i))
}
