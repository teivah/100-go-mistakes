package main

import "fmt"

func main() {
	var n float32 = 1.0001
	fmt.Println(n * n)
}

func f1(n int) float64 {
	result := 10_000.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result
}

func f2(n int) float64 {
	result := 0.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result + 10_000.
}
