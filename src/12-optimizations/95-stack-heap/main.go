package main

func listing1() {
	a := 3
	b := 2

	c := sumValue(a, b)
	println(c)
}

//go:noinline
func sumValue(x, y int) int {
	z := x + y
	return z
}

func listing2() {
	a := 3
	b := 2

	c := sumPtr(a, b)
	println(*c)
}

//go:noinline
func sumPtr(x, y int) *int {
	z := x + y
	return &z
}

func listing3() {
	a := 3
	b := 2
	c := sum(&a, &b)
	println(c)
}

//go:noinline
func sum(x, y *int) int {
	return *x + *y
}
