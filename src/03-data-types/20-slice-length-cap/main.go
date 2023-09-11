package main

import "fmt"

func main() {
	s := make([]int, 3, 6)
	print(s)

	s[1] = 1
	print(s)

	s = append(s, 2)
	print(s)

	s = append(s, 3, 4, 5)
	print(s)

	s1 := make([]int, 3, 6)
	s2 := s1[1:3]
	s1[1] = 1
	print(s2)

	s2 = append(s2, 2)
	print(s1)
	print(s2)

	s2 = append(s2, 3)
	s2 = append(s2, 4)
	s2 = append(s2, 5)
	print(s1)
	print(s2)
}

func print(s []int) {
	fmt.Printf("len=%d, cap=%d: %v\n", len(s), cap(s), s)
}
