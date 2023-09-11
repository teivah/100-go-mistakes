package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	s1 := "Hello, World!"
	s2 := s1[:5]
	fmt.Println(s2)

	s1 = "Hêllo, World!"
	s2 = string([]rune(s1)[:5])
	fmt.Println(s2)
}

type store struct{}

func (s store) handleLog1(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	uuid := log[:36]
	s.store(uuid)
	// Do something
	return nil
}

func (s store) handleLog2(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	uuid := string([]byte(log[:36]))
	s.store(uuid)
	// Do something
	return nil
}

func (s store) handleLog3(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	uuid := string(strings.Clone(log[:36]))
	s.store(uuid)
	// Do something
	return nil
}

func (s store) store(uuid string) {
	// ...
}
