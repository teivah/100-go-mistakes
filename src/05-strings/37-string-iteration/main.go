package main

import "fmt"

func main() {
	s := "hêllo"
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("len=%d\n", len(s))

	for i, r := range s {
		fmt.Printf("position %d: %c\n", i, r)
	}

	runes := []rune(s)
	for i, r := range runes {
		fmt.Printf("position %d: %c\n", i, r)
	}

	s2 := "hello"
	fmt.Printf("%c\n", rune(s2[4]))
}

func getIthRune(largeString string, i int) rune {
	for idx, v := range largeString {
		if idx == i {
			return v
		}
	}
	return -1
}
