package main

import "fmt"

func main() {
	s := &Struct{id: "foo"}
	defer s.print()
	s.id = "bar"
}

type Struct struct {
	id string
}

func (s *Struct) print() {
	fmt.Println(s.id)
}
