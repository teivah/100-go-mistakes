package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s1 []float32
	customer1 := customer{
		ID:         "foo",
		Operations: s1,
	}
	b, _ := json.Marshal(customer1)
	fmt.Println(string(b))

	s2 := make([]float32, 0)
	customer2 := customer{
		ID:         "bar",
		Operations: s2,
	}
	b, _ = json.Marshal(customer2)
	fmt.Println(string(b))
}

type customer struct {
	ID         string
	Operations []float32
}
