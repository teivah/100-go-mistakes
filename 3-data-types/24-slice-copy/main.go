package main

import "fmt"

func bad() {
	src := []int{0, 1, 2}
	var dst []int
	copy(dst, src)
	fmt.Println(dst)

	_ = src
	_ = dst
}

func correct() {
	src := []int{0, 1, 2}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println(dst)

	_ = src
	_ = dst
}
