package main

import "fmt"

func listing1() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}

	fmt.Println(m)
}

func listing2() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}
	m2 := copyMap(m)

	for k, v := range m {
		m2[k] = v
		if v {
			m2[10+k] = true
		}
	}

	fmt.Println(m2)
}

func copyMap(m map[int]bool) map[int]bool {
	res := make(map[int]bool, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}
