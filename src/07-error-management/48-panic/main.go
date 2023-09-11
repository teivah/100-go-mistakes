package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover", r)
		}
	}()

	f()
}

func f() {
	fmt.Println("a")
	panic("foo")
	fmt.Println("b")
}

func checkWriteHeaderCode(code int) {
	if code < 100 || code > 999 {
		panic(fmt.Sprintf("invalid WriteHeader code %v", code))
	}
}
