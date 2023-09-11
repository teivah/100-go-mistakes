package main

import "fmt"

func main() {
	listing1()
	listing2()
	listing3()
}

func listing1() {
	s := []int{1, 2, 3}

	f(s[:2])
	fmt.Println(s)
}

func listing2() {
	s := []int{1, 2, 3}
	sCopy := make([]int, 2)
	copy(sCopy, s)

	f(sCopy)
	result := append(sCopy, s[2])
	fmt.Println(result)
}

func listing3() {
	s := []int{1, 2, 3}
	f(s[:2:2])
	fmt.Println(s)
}

func f(s []int) {
	_ = append(s, 10)
}
