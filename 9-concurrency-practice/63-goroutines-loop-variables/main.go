package main

import "fmt"

func listing1() {
	s := []int{1, 2, 3}

	for _, i := range s {
		go func() {
			fmt.Print(i)
		}()
	}
}

func listing2() {
	s := []int{1, 2, 3}

	for _, i := range s {
		val := i
		go func() {
			fmt.Print(val)
		}()
	}
}

func listing3() {
	s := []int{1, 2, 3}

	for _, i := range s {
		go func(i int) {
			fmt.Print(i)
		}(i)
	}
}
